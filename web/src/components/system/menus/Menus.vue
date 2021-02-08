<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addMenu('0')" type="primary">新增根菜单</el-button>
    </div>

    <el-table :data="tableData"  row-key="id" stripe>
      <el-table-column label="id" min-width="100" prop="id"></el-table-column>
      <el-table-column label="路由Name" min-width="160" prop="name"></el-table-column>
      <el-table-column label="路由Path" min-width="160" prop="path"></el-table-column>
      <el-table-column label="是否隐藏" min-width="100" prop="is_hidden">
        <template slot-scope="scope">
          <span>{{scope.row.is_hidden?"隐藏":"显示"}}</span>
        </template>
      </el-table-column>
      <el-table-column label="父节点" min-width="90" prop="parent_id"></el-table-column>
      <el-table-column label="排序" min-width="70" prop="sort"></el-table-column>
      <el-table-column label="文件路径" min-width="360" prop="component"></el-table-column>
      <el-table-column label="展示名称" min-width="120">
        <template slot-scope="scope">
          <span>{{scope.row.meta.title}}</span>
        </template>
      </el-table-column>
      <el-table-column label="图标" min-width="140">
        <template slot-scope="scope">
          <i :class="`el-icon-${scope.row.meta.icon}`"></i>
          <span>{{scope.row.meta.icon}}</span>
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" width="300">
        <template slot-scope="scope">
          <el-button
              @click="addMenu(scope.row.id)"
              size="small"
              type="primary"
              icon="el-icon-edit"
          >添加子菜单</el-button>
          <el-button
              @click="editMenu(scope.row.id)"
              size="small"
              type="primary"
              icon="el-icon-edit"
          >编辑</el-button>
          <el-button
              @click="deleteMenu(scope.row.id)"
              size="small"
              type="danger"
              icon="el-icon-delete"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>

<!--    弹窗-->
    <el-dialog :before-close="handleClose" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form
          :inline="true"
          :model="form"
          :rules="rules"
          label-position="top"
          label-width="85px"
          ref="menuForm"
      >
        <el-form-item label="菜单名称（中文）" prop="meta.title" style="width:30%">
          <el-input autocomplete="off" v-model="form.meta.title"></el-input>
        </el-form-item>
        <el-form-item label="路由name(path)" prop="name" style="width:30%">
          <el-input
              @change="changeName"
              autocomplete="off"
              placeholder="唯一英文字符串"
              v-model="form.name"
          ></el-input>
        </el-form-item>
        <el-form-item label="文件路径" prop="component" style="width:30%">
          <el-input autocomplete="off" v-model="form.component"></el-input>
        </el-form-item>

        <el-form-item label="父节点Id" prop="parent_id" style="width:30%">
          <el-cascader
              :disabled="!this.isEdit"
              :options="menuOption"
              :props="{ checkStrictly: true,label:'title',value:'id',disabled:'disabled',emitPath:false}"
              :show-all-levels="false"
              filterable
              v-model="form.parent_id"
          ></el-cascader>
        </el-form-item>
        <el-form-item label="是否隐藏" prop="is_hidden" style="width:30%">
          <el-select placeholder="是否在列表隐藏" v-model="form.is_hidden">
            <el-option :value="false" label="否"></el-option>
            <el-option :value="true" label="是"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="排序" prop="sort" style="width:30%">
          <el-input-number v-model.number="form.sort" :min="1" :max="9999"></el-input-number>
        </el-form-item>

        <el-form-item label="图标" prop="meta.icon" style="width:30%">
          <Icon :meta="form.meta">
            <template slot="prepend">el-icon-</template>
          </Icon>
        </el-form-item>
      </el-form>

      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getMenusServer, createMenusServer, getMenuByIdServer, delMenuByIdServer } from '@/api/menu'
import pageData from '@/utils/pageData'
import Icon from './Icon'

export default {
  name: "Menus",
  mixins: [pageData],
  components: {
    Icon
  },
  data () {
    return {
      listApi: getMenusServer,//数据接口
      dialogTitle: '新增菜单',
      dialogFormVisible: false,
      isEdit: false,
      checkFlag: false,
      menuOption: [
        {
          id: "0",
          title: "根菜单"
        }
      ],
      form: {
        id: "0",
        path: "",
        name: "",
        is_hidden: false,
        parent_id: "0",
        component: "",
        sort: 1,
        meta: {
          title: "",
          icon: "",
        },
      },
      rules: {
        name: [{ required: true, message: "请输入菜单name(path)", trigger: "blur" }],
        component: [{ required: true, message: "请输入vue组件路径", trigger: "blur" }],
        is_hidden: [{ required: true, message: "请选择是否隐藏菜单", trigger: "change" }],
        sort: [{ required: true, message: "请设置排序", trigger: "change" }],
        "meta.title": [
          { required: true, message: "请输入菜单中文名称", trigger: "blur" }
        ],
        "meta.icon": [
          { required: true, message: "请输入菜单图标", trigger: "change" }
        ],

      }
    }
  },
  created() {
    this.pageSize = 999;
    this.getTableData();
  },
  methods: {
    handleClose(done) {
      done()
    },
    changeName () {
      //规则：vue的name和path需要一样
      this.form.path = this.form.name
    },
    setOptions() {
      this.menuOption = [
        {
          id: "0",
          title: "根目录"
        }
      ];
      this.setMenuOptions(this.tableData, this.menuOption, false);
    },
    setMenuOptions(menuData, optionsData, disabled) {
      menuData &&
      menuData.map(item => {
        if (item.children && item.children.length) {
          const option = {
            title: item.meta.title,
            id: String(item.id),
            disabled: disabled || item.id == this.form.id,
            children: []
          };
          this.setMenuOptions(
              item.children,
              option.children,
              disabled || item.id == this.form.id
          );
          optionsData.push(option);
        } else {
          const option = {
            title: item.meta.title,
            id: String(item.id),
            disabled: disabled || item.id == this.form.id
          };
          optionsData.push(option);
        }
      });
    },
    addMenu(id) {
      this.dialogTitle = "新增菜单";
      this.form.parent_id = String(id);
      this.isEdit = false;
      this.setOptions();
      this.dialogFormVisible = true
      console.log(id)
    },
    async editMenu(id) {
      this.dialogTitle = "编辑菜单";
      const res = await getMenuByIdServer(id);
      console.log(res)
      this.form = res.data;
      this.form.parent_id = String(this.form.parent_id)
      this.isEdit = true;
      this.setOptions();
      this.dialogFormVisible = true;
    },
    async deleteMenu(id) {
      this.$confirm("此操作将永久删除所有角色下的该菜单", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(async () => {
        const res = await delMenuByIdServer(id)
        if (res.code == 200) {
          this.$message({
            type: "success",
            message: res.msg
          });
          this.getTableData()
        }
      }).catch(() => {
        this.$message({
          type: "info",
          message: "已取消删除"
        });
      })
    },
    closeDialog() {
      this.dialogFormVisible = false
      this.resetForm()
    },
    enterDialog() {
      this.$refs.menuForm.validate(async valid => {
        if (valid) {
          this.form.parent_id = parseInt(this.form.parent_id)
          const res = await createMenusServer(this.form)
          if (res.code == 200) {
            this.$message({
              type: "success",
              message: this.isEdit ? "编辑成功" : "添加成功!"
            });
            this.getTableData();
          }
          this.resetForm();
          this.dialogFormVisible = false;
        }
      })
    },
    resetForm () {
      this.checkFlag = false
      this.$refs.menuForm.resetFields()
      this.form = {
        id: "0",
        path: "",
        name: "",
        is_hidden: false,
        parent_id: "0",
        component: "",
        sort: 1,
        meta: {
          title: "",
          icon: "",
        },
      };


    },
  }
}
</script>

<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.warning {
  color: #dc143c;
}
</style>
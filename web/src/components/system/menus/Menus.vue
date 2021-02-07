<template>
  <div>
    <div class="button-box clearflex">
      <el-button @click="addMenu('0')" type="primary">新增根菜单</el-button>
    </div>

    <el-table :data="tableData"  row-key="id" stripe>
      <el-table-column label="id" min-width="100" prop="id"></el-table-column>
      <el-table-column label="路由Name" min-width="160" prop="name"></el-table-column>
      <el-table-column label="路由Path" min-width="160" prop="path"></el-table-column>
      <el-table-column label="是否隐藏" min-width="100" prop="hidden">
        <template slot-scope="scope">
          <span>{{scope.row.hidden?"隐藏":"显示"}}</span>
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
  </div>
</template>

<script>
import { getMenusServer } from '@/api/menu'
export default {
  name: "Menus",
  data () {
    return {
      tableData: []
    }
  },
  created() {
    this.getMenuList()
  },
  methods: {
    async getMenuList() {
      const res = await getMenusServer()
      this.tableData = res.data
    },
    addMenu(id) {
      console.log(id)
    },
    editMenu(id) {
      console.log(id)
    },
    deleteMenu(id) {
      console.log(id)
    }
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
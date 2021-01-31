<template>
  <div id="userLayout" class="user-layout-wrapper">
    <div class="container">
      <div class="top">
        <div class="desc">
          <img class="logo_login" src="@/assets/logo.png" alt="" />
        </div>
        <div class="header">
          <a href="#">
            <span class="title">仓储管理系统-WMS</span>
          </a>
        </div>
      </div>
<!--      top-->
      <div class="main">
        <el-form
            :model="loginForm"
            :rules="rules"
            ref="loginForm"
        >
          <el-form-item prop="username">
            <el-input placeholder="请输入用户名" v-model="loginForm.username">
              <i class="el-input__icon el-icon-user" /></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input
                type="password"
                placeholder="请输入密码"
                v-model="loginForm.password"
            >
              <i
                  :class="'el-input__icon el-icon-lock'"
              />
            </el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="submitForm" style="width: 100%"
            >登 录</el-button
            >
          </el-form-item>
        </el-form>

      </div>
      <!--/main-->
      <div class="footer">
        <div class="links">
          <a href="#">倚天屠龙记</a>
        </div>
        <div class="copyright">Copyright &copy; </div>
      </div>
<!--      footer-->
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  name: "Login",
  lock: "lock",
  data() {
    return {
      loginForm : {
        username: '',
        password: ''
      },
      rules: {
        username: [
          {required: true, message: '请输入用户名', trigger: "blur"},
          {min: 3, max:12, message: '用户名长度在 3 到 12 个字符', trigger: "blur"}
        ],
        password: [
          {required: true, message: '请输入用密码', trigger: "blur"},
          {min: 6, max:12, message: '密码长度在 6 到 12 个字符', trigger: "blur"}
        ],
      }
    };
  },
  methods: {
    ...mapActions('user', ['LoginAction']),
    async login () {
      return await this.LoginAction(this.loginForm)
    },
    async submitForm() {
      this.$refs.loginForm.validate(async (valid) => {
        if (!valid) {
          this.$message({
            type: 'error',
            message: "请填写正确的的登录信息",
            showClose: false
          });
          return false;
        } else {
          const  flag = await this.login()
          console.log(flag)
        }
      })
    }
  }
}
</script>

<style scoped lang="scss">
@import "@/style/login.scss";
</style>
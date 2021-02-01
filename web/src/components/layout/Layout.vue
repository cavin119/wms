<template>
  <el-container class="layout-cont">
    <el-container class="openside">
      <el-row class="shadowBg"></el-row>
      <el-aside class="main-cont main-left">
        <div class="tilte">
          <img alt class="logoimg" src="~@/assets/nav_logo.png" />
          <h2 class="tit-text">仓储管理-WMS</h2>
        </div>
        <Aside class="aside" />
      </el-aside>
<!--      end el-aside-->

      <el-main class="main-cont main-right">
        <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
          <div
              :style="{width: `calc(100% - ${isMobile?'0px':isCollapse?'54px':'220px'})`}"
              class="topfix"
          >
            <el-row>
              <!-- :xs="8" :sm="6" :md="4" :lg="3" :xl="1" -->
              <el-header class="header-cont">
                <el-col :xs="2" :lg="1" :md="1" :sm="1" :xl="1">
                  <div @click="totalCollapse" class="menu-total">
                    <i class="el-icon-s-unfold" v-if="isCollapse"></i>
                    <i class="el-icon-s-fold" v-else></i>
                  </div>
                </el-col>
                <el-col :xs="10" :lg="14" :md='14' :sm="9" :xl="14">
                  <el-breadcrumb class="breadcrumb" separator-class="el-icon-arrow-right">
                    <el-breadcrumb-item
                        :key="item.path"
                        v-for="item in matched.slice(1,matched.length)"
                    >{{item.meta.title}}</el-breadcrumb-item>
                  </el-breadcrumb>
                </el-col>
                <el-col :xs="12" :lg="9" :md="9" :sm="14" :xl="9">
                  <div class="fl-right right-box">
<!--                    <Search />-->
<!--                    <Screenfull class="screenfull" :style="{cursor:'pointer'}"></Screenfull>-->
                    <el-dropdown>
                  <span class="header-avatar">
<!--                   <CustomPic/>-->
                    <span style="margin-left: 5px">{{userInfo.username}}</span>
                    <i class="el-icon-arrow-down"></i>
                  </span>
                      <el-dropdown-menu class="dropdown-group" slot="dropdown">
                        <el-dropdown-item>
                      <span>
                        更多信息
                        <el-badge is-dot />
                      </span>
                        </el-dropdown-item>
                        <el-dropdown-item icon="el-icon-s-custom">个人信息</el-dropdown-item>
                        <el-dropdown-item icon="el-icon-table-lamp">登 出</el-dropdown-item>
                      </el-dropdown-menu>
                    </el-dropdown>
                  </div>
                </el-col>

              </el-header>
            </el-row>
          </div>
        </transition>
        <transition mode="out-in" name="el-fade-in-linear">
          <keep-alive>
            <router-view :key="$route.fullPath" v-loading="loadingFlag" element-loading-text="正在加载中" class="admin-box"></router-view>
          </keep-alive>
        </transition>
        <Bottom />
      </el-main>
<!--      end main-right-->
    </el-container>
  </el-container>
</template>

<script>
import { mapGetters } from 'vuex'
import Bottom from "./bottom/Bottom";
import Aside from './aside/Aside'

export default {
  name: "Layout",
  components: {
    Aside,
    Bottom
  },
  data() {
    return {
      isMobile: false,
      isCollapse: false,
      isSider: true,
      loadingFlag:false,
      isShadowBg: false
    }
  },//data
  methods: {
    totalCollapse() {
      this.isCollapse = !this.isCollapse
      this.isSider = !this.isCollapse
      this.isShadowBg = !this.isCollapse
      this.$bus.emit('collapse', this.isCollapse)
    }
  },//methods
  computed: {
    ...mapGetters('user', ['userInfo']),
    title() {
      return this.$route.meta.title || '当前页面'
    },
    matched() {
      return this.$route.matched
    }
  },//computed
  mounted() {
    this.$bus.on("showLoading",()=>{
      this.loadingFlag = true
    })
    this.$bus.on("closeLoading",()=>{
      this.loadingFlag = false
    })
  }//mounted

}
</script>
<style lang="scss">
@import '@/style/mobile.scss';
</style>
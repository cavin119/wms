<template>
  <div>
    <el-scrollbar style="height:calc(100vh - 64px)">
      <transition :duration="{ enter: 800, leave: 100 }" mode="out-in" name="el-fade-in-linear">
        <el-menu
            :collapse="isCollapse"
            :collapse-transition="true"
            :default-active="active"
            @select="selectMenuItem"
            active-text-color="#fff"
            class="el-menu-vertical"
            text-color="rgb(191, 203, 217)"
            unique-opened
        >
          <template v-for="item in asyncRouters[0].children">
            <aside-component :key="item.name" :routerInfo="item" v-if="!item.hidden" />
          </template>
        </el-menu>
      </transition>
    </el-scrollbar>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import AsideComponent from "./asideComponent/AsideComponent";

export default {
  name: "Aside",
  data() {
    return {
      active: "",
      isCollapse: false
    };
  },
  components: {
    AsideComponent
  },
  methods: {
    selectMenuItem(index, _, ele) {
      const query = {};
      const params = {};
      ele.route.parameters &&
      ele.route.parameters.map(item => {
        if (item.type == "query") {
          query[item.key] = item.value;
        } else {
          params[item.key] = item.value;
        }
      });
      if (index === this.$route.name) return;
      if (index.indexOf("http://") > -1 || index.indexOf("https://") > -1) {
        window.open(index);
      } else {
        this.$router.push({ name: index, query, params });
      }
    }
  },
  computed: {
    ...mapGetters("router", ["asyncRouters"])
  }
}
</script>

<style scoped>

</style>
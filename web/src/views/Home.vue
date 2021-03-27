<template>
  <div class="init">
    <p class="in-one">欢迎使用初始化工具</p>
    <div class="form-card in-three a-fadeinB">
      <el-form ref="form" :model="form" label-width="100px">
        <el-form-item label="框架类型">
          <el-select v-model="form.frame_type" placeholder="请选择">
            <el-option
              key="gin-vue-admin"
              label="gin-vue-admin"
              value="gin-vue-admin"
            ></el-option>
            <el-option
              key="gf-vue-admin"
              label="gf-vue-admin"
              value="gf-vue-admin"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="语言">
          <el-select v-model="form.language" placeholder="请选择">
            <el-option key="zh-CN" label="简体中文" value="zh-CN"></el-option>
            <el-option key="en" label="English" value="en"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="数据库类型">
          <el-select disabled v-model="form.db_type" placeholder="请选择">
            <el-option
              key="mysql"
              label="mysql(目前只支持mysql)"
              value="mysql"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="host">
          <el-input
            v-model="form.host"
            placeholder="请输入数据库链接"
          ></el-input>
        </el-form-item>
        <el-form-item label="port">
          <el-input
            v-model="form.port"
            placeholder="请输入数据库端口"
          ></el-input>
        </el-form-item>
        <el-form-item label="userName">
          <el-input
            v-model="form.username"
            placeholder="请输入数据库用户名"
          ></el-input>
        </el-form-item>
        <el-form-item label="password">
          <el-input
            v-model="form.password"
            placeholder="请输入数据库密码（没有则为空）"
          ></el-input>
        </el-form-item>
        <el-form-item label="dbName">
          <el-input
            v-model="form.db_name"
            placeholder="请输入数据库名称"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <div style="text-align: right">
            <el-button :disabled="disabled" type="primary" @click.native="onSubmit">立即初始化</el-button>
          </div>
        </el-form-item>
      </el-form>
    </div>
    <p class="in-two a-fadeinT">本工具由flipped-aurora提供，仅用于学习使用！</p>
  </div>
</template>

<script>
// import axios from 'axios' // 引入axios
import { Message } from 'element-ui'

export default {
  name: 'Home',
  data () {
    return {
      form: {
        db_type: 'mysql',
        language: 'zh-CN',
        frame_type: 'gin-vue-admin',
        host: '127.0.0.1',
        port: '3306',
        username: 'root',
        password: '',
        db_name: 'gin_vue_admin'
      },
      disabled: false
    }
  },
  methods: {
    async onSubmit () {
      this.disabled = true
      /* eslint-disable no-undef */
      const value = await submit(JSON.stringify(this.form))
      Message({
        message: value ? '初始化成功' : '初始化失败',
        type: value ? 'success' : 'error'
      })
      this.disabled = false
    }
  }
}
</script>

<style scoped>
.init {
  /* height: 100vh; */
  flex-direction: column;
  display: flex;
  align-items: center;
  background: #fff;
}
.in-one {
  font-size: 26px;
  line-height: 30px;
}
.in-two {
  font-size: 22px;
}
.form-card {
  margin-top: 30px;
  margin-bottom: 30px;
  box-shadow: 0px 0px 5px 0px rgba(5, 12, 66, 0.15);
  width: 60vw;
  height: auto;
  background: #fff;
  padding: 20px;
  border-radius: 6px;
}
</style>

<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model class="loginForm" :model="formdata" :rules="rules">
        <a-form-model-item prop="username">
          <a-input v-model="formdata.username" placeholder="Username">
            <a-icon slot="prefix" type="user" />
          </a-input>
        </a-form-model-item>
        <a-form-model-item prop="password">
          <a-input v-model="formdata.password" placeholder="Password" type="password">
            <a-icon slot="prefix" type="lock" />
          </a-input>
        </a-form-model-item>
        <a-form-model-item class="loginBtn">
          <a-button type="primary" @click="submitForm('ruleForm')" style="margin:10px">
            Submit
          </a-button>
          <a-button type="danger" @click="resetForm('ruleForm')" style="margin:10px">
            Reset
          </a-button>
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
export default {
  data () {
    return {
      formdata: {
        username: '',
        password: ''
      },
      rules: {
        username: [
          { required: true, message: '請輸入用戶名', trigger: 'blur' },
          { min: 4, max: 12, message: '用戶名長度在4到12個字符之間', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '請輸入密碼', trigger: 'blur' },
          { min: 6, max: 12, message: '密碼長度在6到12個字符之間', trigger: 'blur' }
        ]
      }
    }
  }
}
</script>

<style scoped>
.container {
  height: 100%;
  background-color: #282c34;
}

.loginBox {
  width: 450px;
  height: 300px;
  background-color: #fff;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}
.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
}
.loginBtn {
  display: flex;
  justify-content: flex-end;
}
</style>

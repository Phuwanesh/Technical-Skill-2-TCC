<template>
  <div class="container">
    <div class="header">
      <h3>IT 02-1</h3>
    </div>

    <div class="form-box">
      <form @submit.prevent="login" class="form">
        <div class="form-group">
          <label>User</label>
          <input v-model="username" type="text" required />
        </div>

        <div class="form-group">
          <label>Password</label>
          <input v-model="password" type="password" required />
        </div>

        <div class="form-group">
          <label></label>
          <button type="submit" class="login-btn">ลงชื่อเข้าใช้งาน</button>
        </div>

        <div class="form-group">
          <label></label>
          <router-link to="/register" class="register-link">
            สมัครสมาชิก
          </router-link>        
        </div>

        <p v-if="errorMessage" class="error-msg">{{ errorMessage }}</p>
      </form>


    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import api from "../api";

const router = useRouter();
const username = ref("");
const password = ref("");
const errorMessage = ref("");

const login = async () => {
  errorMessage.value = "";
  try {
    const res = await api.post("/login", {
      username: username.value,
      password: password.value,
    });
    localStorage.setItem("token", res.data.token);
    router.push("/home");
  } catch (err) {
    if (err.response?.data?.error) {
      errorMessage.value = err.response.data.error;
    } else {
      errorMessage.value = "เข้าสู่ระบบไม่สำเร็จ กรุณาลองใหม่อีกครั้ง";
    }
  }
};
</script>

<style scoped>
html,
body,
#app {
  margin: 0;
  padding: 0;
  height: 100%;
}


.login-btn {
  padding: 10px 25px;
  border: 1px solid #009344;
  background: #fff;
  cursor: pointer;
  border-radius: 5px;
  font-size: 16px;
  height: 50px;
}

.register-link {
  display: block;
  color: blue;
  font-size: 14px;
  text-decoration: underline;
}

@media (max-width: 600px) {

.form-group {
    flex-direction: column;
    align-items: flex-start;  
    width: 100%;
  }

.form-group label {
    text-align: left;         
    margin: 0 0 5px 0;        
    width: 100%;              
    font-weight: bold;
  }
}
</style>

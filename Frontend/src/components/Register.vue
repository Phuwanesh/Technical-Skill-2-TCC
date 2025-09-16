<template>
  <div class="container">
    <div class="header">
      <h3>IT 02-2</h3>
    </div>

    <div class="form-box">
      <form @submit.prevent="register" class="form">
        <div class="form-group">
          <label>User</label>
          <input v-model="username" type="text" required />
        </div>

        <div class="form-group">
          <label>Password</label>
          <input v-model="password" type="password" required />
        </div>

        <div class="form-group">
          <label>Confirm Password</label>
          <input v-model="confirmPassword" type="password" required />
        </div>

        <button type="submit" class="register-btn">สมัครสมาชิก</button>

        <p v-if="errorMessage" class="error-msg">{{ errorMessage }}</p>
      </form>

      <!-- <router-link to="/" class="login-link">
        กลับไปหน้าเข้าสู่ระบบ
      </router-link> -->
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
const confirmPassword = ref("");
const errorMessage = ref("");

const register = async () => {
  errorMessage.value = "";
  try {
    await api.post("/register", {
      username: username.value,
      password: password.value,
      confirmPassword: confirmPassword.value,
    });
    alert("สมัครสมาชิกสำเร็จ");
    router.push("/");
  } catch (err) {
    if (err.response?.data?.error) {
      errorMessage.value = err.response.data.error;
    } else {
      errorMessage.value = "สมัครสมาชิกไม่สำเร็จ กรุณาลองใหม่อีกครั้ง";
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

.register-btn {
  margin-top: 20px;
  padding: 10px 25px;
  border: 1px solid #009344;
  background: #fff;
  cursor: pointer;
  border-radius: 5px;
  font-size: 16px;
}

.login-link {
  display: block;
  margin-top: 15px;
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

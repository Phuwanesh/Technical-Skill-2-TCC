<template>
  <div class="container">
    <div class="header">
      <h3>IT 02-3</h3>
    </div>
    <div class="form-box">
      <p v-if="user">Welcome User : {{ user.username }}</p>
      <!-- ปุ่ม Logout -->
      <!-- <button @click="logout" class="logout-btn">Logout</button> -->
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import api from "../api";

const router = useRouter();
const user = ref(null);

onMounted(async () => {
  try {
    const res = await api.get("/me");
    user.value = res.data;
  } catch {
    user.value = null;
  }
});

const logout = () => {
  localStorage.removeItem("token"); 
  router.push("/");                 
};
</script>

<style scoped>
.logout-btn {
  margin-top: 20px;
  padding: 8px 16px;
  border: 1px solid #d9534f;
  background: #d9534f;
  color: white;
  cursor: pointer;
  border-radius: 5px;
  font-size: 14px;
}

.logout-btn:hover {
  background: #c9302c;
}
</style>

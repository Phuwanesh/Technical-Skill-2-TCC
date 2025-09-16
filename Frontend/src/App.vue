<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script setup>
</script>

<style>
#app {
  font-family: Arial, sans-serif;
}
  .container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  background: #fff;
  border: 4px solid #000;
  font-size: 18px;
}

  .header {
  background: #009344;
  color: black;
  padding: 15px;
  text-align: center;
  border-bottom: 4px solid black;
}

.form {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 120%;
}

.form-box {
  display: flex;
  flex-direction: column;
  align-items: center;   
  justify-content: center;
  height: 70%;
}

.form-group {
  margin: 15px 0;
  display: flex;
  align-items: center;
  width: 100%;
  max-width: 500px;
}

.form-group label {
  width: 170px;
  text-align: right;
  margin-right: 10px;
  font-weight: bold;
  white-space: nowrap;
}

.form-group input {
  flex: 1;
  padding: 10px;
  border: 1px solid #333;
  font-size: 16px;
}

.form-group.with-button {
  gap: 10px;
}

.with-button .login-btn {
  padding: 10px 20px;
  border: 1px solid #009344;
  background: #fff;
  cursor: pointer;
  border-radius: 5px;
  font-size: 16px;
  white-space: nowrap;
}

.error-msg {
  margin-top: 15px;
  color: red;
  font-size: 14px;
  text-align: center;
}

@media (max-width: 600px) {
  .form {
    width: 100%;
    padding: 0 20px;          
    box-sizing: border-box;
  }

  .form-group input {
    width: 100%;              
    box-sizing: border-box;
  }

  .login-btn {
    width: 100%;
    max-width: none;
    margin-top: 20px;
  }
}
</style>

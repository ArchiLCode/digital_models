<template>
  <div class="container">
    <h2>Contibutors</h2>
    <div class="contributors">
      <img src="@/assets/contributors/dior.png" alt="" />
      <img src="@/assets/contributors/nike.png" alt="" />
      <img src="@/assets/contributors/gucci.png" alt="" />
      <img src="@/assets/contributors/lv.png" alt="" />
      <img src="@/assets/contributors/adidas.png" alt="" />
      <img src="@/assets/contributors/zara.png" alt="" />
      <img src="@/assets/contributors/chanel.png" alt="" />
      <img src="@/assets/contributors/uniqlo.png" alt="" />
      <img src="@/assets/contributors/cartier.png" alt="" />
      <img src="@/assets/contributors/hm.png" alt="" />
      <img src="@/assets/contributors/tiffany.png" alt="" />
      <img src="@/assets/contributors/prada.png" alt="" />
      <img src="@/assets/contributors/puma.png" alt="" />
      <img src="@/assets/contributors/fendi.png" alt="" />
      <img src="@/assets/contributors/maxmara.png" alt="" />
      <img src="@/assets/contributors/calvin-klein.png" alt="" />
      <img src="@/assets/contributors/vogue.png" alt="" />
    </div>
  </div>
</template>
<script setup></script>
<style scoped>
h2 {
  font-size: 16px;
  margin-top: 80px;
  text-align: center;
}
img {
  height: 25px;
  transition: 0.1s;
}

img:hover {
  scale: 1.2;
}
.contributors {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: center;
  margin: 30px 0;
  gap: 35px;
}
.container {
  margin: 0 auto;
  max-width: 800px;
  width: 100%;
}
</style>

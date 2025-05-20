<template>
  <footer class="footer">
    <p class="copyright">
      &copy; 2013-2025, Model management ÉLEFÉ Inc. All rights reserved.
    </p>
    <p class="inspired">Inspired by Max Shnek</p>
  </footer>
</template>

<script setup></script>

<style scoped>
.footer {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  margin-top: 60px;
  padding-bottom: 20px;
}

.footer p {
  flex: 1;
  font-size: 12px;
  font-weight: 600;
  color: #f8f8f8;
}

.footer p:first-child {
  text-align: left;
}

.footer p:last-child {
  text-align: right;
}

@media (max-width: 768px) {
  .footer {
    flex-direction: column;
    align-items: center;
    text-align: center;
    padding: 20px 0;
  }

  .footer p.inspired {
    order: 1;
    text-align: center;
    margin: 10px 0;
  }

  .footer p.copyright {
    order: 2;
    text-align: center;
    margin: 10px 0;
  }
}
</style>

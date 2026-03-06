<template>
  <footer class="footer">
    <div class="container">
      <div class="footer-content">
        <div class="footer-brand">
          <HeartHandshake class="logo-icon" :size="36" />
          <span class="logo-text">萌宠美容工坊</span>
          <p class="footer-desc">用爱心和专业，给每一位毛孩子最温柔的呵护</p>
        </div>
        <div class="footer-links">
          <h4>快速导航</h4>
          <a href="#home">首页</a>
          <a href="#stylings">造型</a>
          <a href="#packages">套餐</a>
          <a href="#appointment">预约</a>
        </div>
        <div class="footer-contact">
          <h4>联系我们</h4>
          <p><Phone :size="16" class="inline-icon" /> 400-888-9999</p>
          <p><Mail :size="16" class="inline-icon" /> hello@petbeauty.com</p>
          <p><MapPin :size="16" class="inline-icon" /> XX市XX区XX路100号</p>
        </div>
      </div>
      <div class="footer-bottom">
        <p>© 2026 萌宠美容工坊 All Rights Reserved.</p>
      </div>
    </div>
  </footer>
</template>

<script setup>
import { HeartHandshake, Phone, Mail, MapPin } from 'lucide-vue-next'
</script>

<style lang="scss" scoped>
.footer {
  background: #0C4A6E; /* 对应 text 主色调作为深色页脚 */
  color: #F8FAFC;
  padding: 80px 0 0;
}

.footer-content {
  display: grid;
  grid-template-columns: 2fr 1fr 1fr;
  gap: 40px;
  margin-bottom: 40px;
}

.footer-brand {
  .logo-icon {
    color: #38BDF8;
    vertical-align: middle;
  }

  .logo-text {
    font-size: 1.5rem;
    font-weight: 800;
    color: white;
    margin-left: 8px;
  }

  .footer-desc {
    margin-top: 12px;
    font-size: 0.9rem;
    line-height: 1.6;
    max-width: 300px;
  }
}

.footer-links, .footer-contact {
  h4 {
    color: white;
    font-size: 1rem;
    font-weight: 700;
    margin-bottom: 16px;
  }
}

.footer-links {
  display: flex;
  flex-direction: column;
  gap: 8px;

  a {
    color: rgba(255, 255, 255, 0.7);
    text-decoration: none;
    font-size: 0.9rem;
    transition: color 0.3s ease;

    &:hover {
      color: #38BDF8;
      padding-left: 4px;
    }
  }
}

.footer-contact p {
  font-size: 0.95rem;
  margin-bottom: 12px;
  display: flex;
  align-items: center;
  gap: 8px;
  color: #CBD5E1;
}

.footer-bottom {
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  padding: 20px 0;
  text-align: center;
  font-size: 0.85rem;
}

@media (max-width: 768px) {
  .footer-content {
    grid-template-columns: 1fr;
    gap: 28px;
  }
}
</style>

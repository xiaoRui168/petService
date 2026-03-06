<template>
  <nav class="navbar" :class="{ scrolled: isScrolled }">
    <div class="navbar-inner container">
      <a href="#home" class="logo">
        <HeartHandshake class="logo-icon" :size="32" />
        <span class="logo-text">萌宠美容工坊</span>
      </a>
      <button class="mobile-toggle" @click="menuOpen = !menuOpen" :class="{ active: menuOpen }">
        <span></span><span></span><span></span>
      </button>
      <ul class="nav-links" :class="{ open: menuOpen }">
        <li v-for="item in navItems" :key="item.href">
          <a :href="item.href" @click="menuOpen = false">{{ item.label }}</a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { HeartHandshake } from 'lucide-vue-next'

const isScrolled = ref(false)
const menuOpen = ref(false)

const navItems = [
  { href: '#home', label: '首页' },
  { href: '#stylings', label: '造型' },
  { href: '#packages', label: '套餐' },
  { href: '#gallery', label: '照片' },
  { href: '#reviews', label: '评价' },
  { href: '#appointment', label: '预约' }
]

const onScroll = () => {
  isScrolled.value = window.scrollY > 50
}

onMounted(() => window.addEventListener('scroll', onScroll))
onUnmounted(() => window.removeEventListener('scroll', onScroll))
</script>

<style lang="scss" scoped>
.navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100;
  padding: 16px 0;
  transition: all 0.3s ease;
  background: transparent;

  &.scrolled {
    background: rgba(255, 255, 255, 0.92);
    backdrop-filter: blur(20px);
    box-shadow: 0 2px 20px rgba(255, 154, 139, 0.12);
    padding: 10px 0;
  }
}

.navbar-inner {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.logo {
  display: flex;
  align-items: center;
  gap: 8px;
  text-decoration: none;
  font-weight: 800;
  font-size: 1.4rem;
  color: var(--color-text, #0C4A6E);

  .logo-icon {
    color: var(--color-primary, #0EA5E9);
    animation: float 3s ease-in-out infinite;
  }
}

.nav-links {
  display: flex;
  list-style: none;
  gap: 8px;

  a {
    text-decoration: none;
    color: var(--color-text-light, #475569);
    font-weight: 600;
    padding: 8px 16px;
    border-radius: 50px;
    transition: all 0.3s ease;
    font-size: 1rem;

    &:hover {
      background: rgba(14, 165, 233, 0.1);
      color: var(--color-primary, #0EA5E9);
    }
  }
}

.mobile-toggle {
  display: none;
  flex-direction: column;
  gap: 5px;
  background: none;
  border: none;
  cursor: pointer;
  padding: 4px;

  span {
    display: block;
    width: 24px;
    height: 2.5px;
    background: #4A3728;
    border-radius: 2px;
    transition: all 0.3s ease;
  }

  &.active span:nth-child(1) { transform: translateY(7.5px) rotate(45deg); }
  &.active span:nth-child(2) { opacity: 0; }
  &.active span:nth-child(3) { transform: translateY(-7.5px) rotate(-45deg); }
}

@media (max-width: 768px) {
  .mobile-toggle { display: flex; }

  .nav-links {
    position: absolute;
    top: 100%;
    left: 0;
    right: 0;
    background: rgba(255, 255, 255, 0.97);
    backdrop-filter: blur(20px);
    flex-direction: column;
    padding: 16px 24px;
    gap: 4px;
    display: none;
    box-shadow: 0 10px 30px rgba(0,0,0,0.08);

    &.open { display: flex; }

    a {
      padding: 12px 16px;
      border-radius: 12px;
    }
  }
}
</style>

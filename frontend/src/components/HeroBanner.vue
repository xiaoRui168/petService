<template>
  <section id="home" class="hero">
    <div class="hero-bg">
      <Sparkles class="floating-shape shape-1" :size="48" stroke-width="1.5" />
      <Heart class="floating-shape shape-2" :size="64" stroke-width="1.5" />
      <div class="floating-shape shape-3 blur-blob blob-blue"></div>
      <div class="floating-shape shape-4 blur-blob blob-orange"></div>
    </div>
    <div class="container hero-content">
      <div class="hero-text">
        <span class="hero-badge"><Sparkles :size="16" class="inline-icon" /> 专业宠物美容</span>
        <h1>给毛孩子<br/><span class="gradient-text">最温柔的呵护</span></h1>
        <p class="hero-desc">我们用爱心和专业技术，让每一位毛茸茸的宝贝都变得更加可爱迷人。从洗护到造型，给您的爱宠最舒适的美容体验。</p>
        <div class="hero-actions">
          <a href="#appointment" @click.prevent="scrollTo('appointment')" class="btn btn-primary"><CalendarPlus :size="20" /> 立即预约</a>
          <a href="#stylings" @click.prevent="scrollTo('stylings')" class="btn btn-outline"><Palette :size="20" /> 查看造型</a>
        </div>
        <div class="hero-stats">
          <div class="stat-item">
            <span class="stat-num">2000+</span>
            <span class="stat-label">服务宝贝</span>
          </div>
          <div class="stat-item">
            <span class="stat-num">50+</span>
            <span class="stat-label">造型方案</span>
          </div>
          <div class="stat-item">
            <span class="stat-num">4.9</span>
            <span class="stat-label">用户评分</span>
          </div>
        </div>
      </div>
      <div class="hero-visual">
        <div class="hero-image-wrapper">
          <img src="https://images.unsplash.com/photo-1587300003388-59208cc962cb?w=600&h=600&fit=crop" alt="可爱的宠物" />
          <div class="hero-image-badge"><Star class="inline-icon text-yellow" :size="16" fill="currentColor" /> Top 美容师推荐</div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { Sparkles, Heart, CalendarPlus, Palette, Star } from 'lucide-vue-next'

const scrollTo = (id) => {
  const el = document.getElementById(id)
  if (el) {
    const offset = 70 // 预留导航条高度
    const top = el.getBoundingClientRect().top + window.scrollY - offset
    window.scrollTo({ top, behavior: 'smooth' })
  }
}
</script>

<style lang="scss" scoped>
.hero {
  position: relative;
  min-height: 100vh;
  display: flex;
  align-items: center;
  overflow: hidden;
  background: linear-gradient(165deg, #F0F9FF 0%, #E0F2FE 40%, #FAFAFA 100%);
}

.hero-bg {
  position: absolute;
  inset: 0;
  overflow: hidden;
}

.floating-shape {
  position: absolute;
  color: #38BDF8;
  opacity: 0.15;
  animation: float 6s ease-in-out infinite;

  &.shape-1 { top: 15%; left: 8%; animation-delay: 0s; color: #F97316; }
  &.shape-2 { top: 25%; right: 12%; animation-delay: 1.2s; }
}

.blur-blob {
  position: absolute;
  border-radius: 50%;
  filter: blur(80px);
  z-index: 0;
  animation: float 10s ease-in-out infinite;

  &.blob-blue {
    width: 400px;
    height: 400px;
    background: rgba(14, 165, 233, 0.2);
    top: -100px;
    left: -100px;
  }
  &.blob-orange {
    width: 300px;
    height: 300px;
    background: rgba(249, 115, 22, 0.15);
    bottom: -50px;
    right: -50px;
    animation-delay: -5s;
  }
}

.hero-content {
  position: relative;
  z-index: 10;
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 60px;
  align-items: center;
  padding-top: 80px;
}

.hero-badge {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 8px 24px;
  background: white;
  border: 2px solid #0EA5E9;
  border-radius: 50px;
  font-size: 0.95rem;
  font-weight: 700;
  color: #0EA5E9;
  margin-bottom: 24px;
  box-shadow: 0 4px 10px rgba(14, 165, 233, 0.1);
}

.hero-text h1 {
  font-size: 4rem;
  font-weight: 800;
  line-height: 1.15;
  margin-bottom: 24px;
  color: #0C4A6E;
  letter-spacing: -1px;
}

.gradient-text {
  background: linear-gradient(135deg, #0EA5E9, #38BDF8);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.hero-desc {
  font-size: 1.25rem;
  color: #475569;
  line-height: 1.7;
  margin-bottom: 40px;
  max-width: 520px;
}

.hero-actions {
  display: flex;
  gap: 16px;
  margin-bottom: 48px;
}

.hero-stats {
  display: flex;
  gap: 40px;

  .stat-item {
    display: flex;
    flex-direction: column;
  }

  .stat-num {
    font-size: 2.2rem;
    font-weight: 800;
    color: #0EA5E9;
  }

  .stat-label {
    font-size: 0.9rem;
    color: #475569;
    font-weight: 600;
    text-transform: uppercase;
    letter-spacing: 0.5px;
  }
}

.hero-visual {
  display: flex;
  justify-content: center;
}

.hero-image-wrapper {
  position: relative;

  img {
    width: 460px;
    height: 460px;
    border-radius: 40px;
    object-fit: cover;
    border: 8px solid white;
    box-shadow: 10px 10px 20px rgba(14, 165, 233, 0.15), -10px -10px 20px rgba(255, 255, 255, 0.8);
    transform: rotate(2deg);
    transition: transform 0.4s ease;

    &:hover {
      transform: rotate(0deg) scale(1.02);
    }
  }

  .hero-image-badge {
    position: absolute;
    bottom: -15px;
    left: 50%;
    transform: translateX(-50%);
    background: white;
    padding: 12px 28px;
    border-radius: 50px;
    font-weight: 700;
    font-size: 1rem;
    color: #0C4A6E;
    box-shadow: 0 10px 25px rgba(14, 165, 233, 0.2);
    white-space: nowrap;
    border: 3px solid #0EA5E9;
    display: flex;
    align-items: center;
    gap: 6px;
  }
}

.text-yellow {
  color: #FACC15;
}

@media (max-width: 768px) {
  .hero-content {
    grid-template-columns: 1fr;
    text-align: center;
    gap: 40px;
  }

  .hero-text h1 { font-size: 2.4rem; }
  .hero-desc { max-width: 100%; }
  .hero-actions { justify-content: center; flex-wrap: wrap; }
  .hero-stats { justify-content: center; }

  .hero-image-wrapper img {
    width: 280px;
    height: 280px;
  }
}
</style>

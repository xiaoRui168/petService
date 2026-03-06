<template>
  <section id="stylings" class="section stylings-section">
    <div class="container">
      <h2 class="section-title"><Palette class="inline-icon" :size="36" stroke-width="2.5" /> 精选造型</h2>
      <p class="section-subtitle">多种风格任你挑选，总有一款适合你的毛孩子</p>
      <div class="stylings-grid">
        <div class="card styling-card" v-for="item in stylings" :key="item.id">
          <div class="card-image">
            <img :src="item.image_url" :alt="item.name" loading="lazy" />
            <div class="card-price">¥{{ item.price }}</div>
          </div>
          <div class="card-body">
            <h3>{{ item.name }}</h3>
            <p>{{ item.description }}</p>
            <a href="#appointment" class="btn btn-primary btn-sm">预约此造型</a>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getStylings } from '../api'
import { Palette } from 'lucide-vue-next'

const stylings = ref([])

onMounted(async () => {
  try {
    const res = await getStylings()
    if (res.data.code === 200) stylings.value = res.data.data
  } catch (e) {
    console.error('获取造型失败', e)
  }
})
</script>

<style lang="scss" scoped>
.stylings-section {
  background: linear-gradient(180deg, #FFF8F0, #FFFFFF);
}

.stylings-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 28px;
}

.styling-card {
  .card-image {
    position: relative;
    height: 240px;
    overflow: hidden;
    border-bottom: 4px solid #F0F9FF;

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
      transition: transform 0.5s ease;
    }

    .card-price {
      position: absolute;
      top: 16px;
      right: 16px;
      background: #0EA5E9;
      color: white;
      padding: 8px 18px;
      border-radius: 50px;
      font-weight: 700;
      font-size: 1rem;
      box-shadow: 0 4px 10px rgba(14, 165, 233, 0.3);
      border: 2px solid white;
    }
  }

  &:hover .card-image img {
    transform: scale(1.05);
  }

  .card-body {
    padding: 24px;

    h3 {
      font-size: 1.35rem;
      font-weight: 700;
      margin-bottom: 8px;
      color: #0C4A6E;
    }

    p {
      color: #475569;
      font-size: 0.95rem;
      line-height: 1.6;
      margin-bottom: 20px;
    }

    .btn-sm {
      padding: 10px 24px;
      font-size: 0.95rem;
    }
  }
}

@media (max-width: 768px) {
  .stylings-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .stylings-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>

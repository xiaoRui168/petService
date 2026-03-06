<template>
  <section id="packages" class="section packages-section">
    <div class="container">
      <h2 class="section-title"><Sparkles class="inline-icon" :size="36" stroke-width="2.5" /> 服务套餐</h2>
      <p class="section-subtitle">从基础洗护到豪华 SPA，总有一款适合你</p>
      <div class="packages-grid">
        <div class="card package-card" v-for="(pkg, index) in packages" :key="pkg.id"
             :class="{ featured: index === 2 }">
          <div class="package-ribbon" v-if="index === 2"><Flame :size="14" class="inline-icon" /> 最受欢迎</div>
          <div class="package-header">
            <div class="package-icon">
               <component :is="icons[index % icons.length]" :size="48" stroke-width="1.5" />
            </div>
            <h3>{{ pkg.name }}</h3>
            <p class="package-desc">{{ pkg.description }}</p>
          </div>
          <div class="package-price">
            <span class="price-currency">¥</span>
            <span class="price-amount">{{ pkg.price }}</span>
            <span class="price-unit">/次</span>
          </div>
          <ul class="package-services">
            <li v-for="(s, i) in pkg.services.split(',')" :key="i">
              <CheckCircle2 class="check" :size="18" /> {{ s }}
            </li>
          </ul>
          <a href="#appointment" class="btn" :class="index === 2 ? 'btn-primary' : 'btn-outline'" style="width: 100%; justify-content: center;">
            立即预约
          </a>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getPackages } from '../api'
import { Sparkles, Flame, CheckCircle2, Bath, Scissors, Crown, Cat } from 'lucide-vue-next'

const packages = ref([])
const icons = [Bath, Scissors, Crown, Cat]

onMounted(async () => {
  try {
    const res = await getPackages()
    if (res.data.code === 200) packages.value = res.data.data
  } catch (e) {
    console.error('获取套餐失败', e)
  }
})
</script>

<style lang="scss" scoped>
.packages-section {
  background: #FFFFFF;
}

.packages-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 24px;
  align-items: start;
}

.package-card {
  padding: 32px 24px;
  text-align: center;
  position: relative;
  border: 4px solid white;

  &.featured {
    border-color: #0EA5E9;
    background: #F8FAFC;
    transform: scale(1.05);
  }

  .package-ribbon {
    position: absolute;
    top: -4px;
    left: 50%;
    transform: translateX(-50%);
    background: #F97316;
    color: white;
    padding: 6px 20px;
    border-radius: 0 0 16px 16px;
    font-size: 0.85rem;
    font-weight: 700;
    box-shadow: 0 4px 10px rgba(249, 115, 22, 0.2);
    display: flex;
    align-items: center;
    gap: 4px;
  }
}

.package-header {
  margin-bottom: 20px;

  .package-icon {
    color: #38BDF8;
    margin-bottom: 16px;
    display: flex;
    justify-content: center;
  }

  h3 {
    font-size: 1.4rem;
    font-weight: 700;
    margin-bottom: 8px;
    color: #0C4A6E;
  }

  .package-desc {
    font-size: 0.95rem;
    color: #475569;
  }
}

.package-price {
  margin-bottom: 24px;
  color: #0EA5E9;

  .price-currency {
    font-size: 1.2rem;
    font-weight: 700;
    vertical-align: top;
    margin-right: 2px;
  }

  .price-amount {
    font-size: 3.2rem;
    font-weight: 800;
    line-height: 1;
    letter-spacing: -1px;
  }

  .price-unit {
    font-size: 0.95rem;
    color: #64748B;
    margin-left: 2px;
  }
}

.package-services {
  list-style: none;
  text-align: left;
  margin-bottom: 32px;

  li {
    padding: 8px 0;
    font-size: 0.95rem;
    color: #0C4A6E;
    font-weight: 500;
    display: flex;
    align-items: center;
    gap: 12px;

    .check {
      color: #0EA5E9;
    }
  }
}

@media (max-width: 768px) {
  .packages-grid {
    grid-template-columns: 1fr;
    gap: 20px;
  }

  .package-card.featured {
    transform: none;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .packages-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>

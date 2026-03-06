<template>
  <section id="reviews" class="section reviews-section">
    <div class="container">
      <h2 class="section-title"><MessageCircleHeart class="inline-icon" :size="36" stroke-width="2.5" /> 宠主好评</h2>
      <p class="section-subtitle">听听其他宠物主人怎么说</p>
      <div class="reviews-slider" ref="sliderRef">
        <div class="reviews-track" :style="{ transform: `translateX(-${currentSlide * slideWidth}px)` }">
          <div class="review-card" v-for="review in reviews" :key="review.id">
            <div class="review-header">
              <div class="review-avatar">{{ review.nickname.charAt(0) }}</div>
              <div class="review-info">
                <span class="review-name">{{ review.nickname }}</span>
                <span class="review-pet">{{ review.pet_type }}</span>
              </div>
              <div class="stars">
                <Star class="star" v-for="i in 5" :key="i" :class="{ filled: i <= review.rating }" :size="16" :fill="i <= review.rating ? 'currentColor' : 'transparent'" stroke-width="2" />
              </div>
            </div>
            <p class="review-content">{{ review.content }}</p>
            <span class="review-date">{{ formatDate(review.created_at) }}</span>
          </div>
        </div>
      </div>
      <div class="slider-controls">
        <button class="slider-btn" @click="prevSlide" :disabled="currentSlide === 0"><ChevronLeft :size="24" /></button>
        <div class="slider-dots">
          <span v-for="i in maxSlides" :key="i" class="dot" :class="{ active: currentSlide === i - 1 }" @click="currentSlide = i - 1"></span>
        </div>
        <button class="slider-btn" @click="nextSlide" :disabled="currentSlide >= maxSlides - 1"><ChevronRight :size="24" /></button>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { getReviews } from '../api'
import { MessageCircleHeart, Star, ChevronLeft, ChevronRight } from 'lucide-vue-next'

const reviews = ref([])
const currentSlide = ref(0)
const sliderRef = ref(null)
const slideWidth = ref(380)

const maxSlides = computed(() => Math.max(1, reviews.value.length - 1))

const prevSlide = () => { if (currentSlide.value > 0) currentSlide.value-- }
const nextSlide = () => { if (currentSlide.value < maxSlides.value - 1) currentSlide.value++ }

const formatDate = (dateStr) => {
  const d = new Date(dateStr)
  return `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}-${String(d.getDate()).padStart(2, '0')}`
}

const updateSlideWidth = () => {
  if (window.innerWidth < 768) slideWidth.value = window.innerWidth - 48
  else slideWidth.value = 380
}

onMounted(async () => {
  updateSlideWidth()
  window.addEventListener('resize', updateSlideWidth)
  try {
    const res = await getReviews()
    if (res.data.code === 200) reviews.value = res.data.data
  } catch (e) {
    console.error('获取评价失败', e)
  }
})

onUnmounted(() => window.removeEventListener('resize', updateSlideWidth))
</script>

<style lang="scss" scoped>
.reviews-section {
  background: linear-gradient(180deg, #F0F9FF, #FFFFFF);
  overflow: hidden;
}

.reviews-slider {
  overflow: hidden;
  margin-bottom: 24px;
}

.reviews-track {
  display: flex;
  gap: 24px;
  transition: transform 0.5s ease;
}

.review-card {
  min-width: 356px;
  background: white;
  border-radius: 20px;
  padding: 28px;
  box-shadow: 8px 8px 16px rgba(14, 165, 233, 0.1), -8px -8px 16px white;
  border: 3px solid #F0F9FF;

  .review-header {
    display: flex;
    align-items: center;
    gap: 12px;
    margin-bottom: 16px;
  }

  .review-avatar {
    width: 48px;
    height: 48px;
    border-radius: 50%;
    background: linear-gradient(135deg, #38BDF8, #0EA5E9);
    display: flex;
    align-items: center;
    justify-content: center;
    color: white;
    font-weight: 800;
    font-size: 1.2rem;
    box-shadow: 0 4px 10px rgba(14, 165, 233, 0.3);
  }

  .review-info {
    flex: 1;
    display: flex;
    flex-direction: column;

    .review-name {
      font-weight: 700;
      color: #0C4A6E;
      font-size: 1rem;
    }

    .review-pet {
      font-size: 0.85rem;
      color: #64748B;
    }
  }

  .review-content {
    color: #475569;
    font-size: 1rem;
    line-height: 1.7;
    margin-bottom: 20px;
  }

  .review-date {
    font-size: 0.8rem;
    color: #aaa;
  }
}

.slider-controls {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 16px;
}

.slider-btn {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  border: none;
  background: white;
  color: #0EA5E9;
  box-shadow: 0 4px 10px rgba(14, 165, 233, 0.2), 0 0 0 2px #0EA5E9 inset;
  cursor: pointer;
  transition: all 0.2s cubic-bezier(0.34, 1.56, 0.64, 1);
  display: flex;
  align-items: center;
  justify-content: center;

  &:hover:not(:disabled) {
    background: #0EA5E9;
    color: white;
    transform: scale(1.1);
  }

  &:active:not(:disabled) {
    transform: scale(0.95);
  }

  &:disabled {
    opacity: 0.4;
    cursor: not-allowed;
    box-shadow: none;
    border: 2px solid #E2E8F0;
    color: #94A3B8;
  }
}

.slider-dots {
  display: flex;
  gap: 8px;

  .dot {
    width: 10px;
    height: 10px;
    border-radius: 50%;
    background: #E2E8F0;
    cursor: pointer;
    transition: all 0.3s ease;

    &.active {
      background: #F97316;
      transform: scale(1.4);
      box-shadow: 0 2px 6px rgba(249, 115, 22, 0.3);
    }
  }
}

@media (max-width: 768px) {
  .review-card {
    min-width: calc(100vw - 48px);
  }
}
</style>

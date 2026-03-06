<template>
  <section id="gallery" class="section gallery-section">
    <div class="container">
      <h2 class="section-title"><Camera class="inline-icon" :size="36" stroke-width="2.5" /> 萌宠画廊</h2>
      <p class="section-subtitle">看看我们的毛孩子们美容后的可爱模样</p>
      <div class="gallery-grid">
        <div class="gallery-item" v-for="photo in photos" :key="photo.id" @click="openLightbox(photo)">
          <img :src="photo.image_url" :alt="photo.pet_name" loading="lazy" />
          <div class="gallery-overlay">
            <span class="pet-name">{{ photo.pet_name }}</span>
            <span class="pet-desc">{{ photo.description }}</span>
            <span class="groomer" v-if="photo.groomer">美容师: {{ photo.groomer }}</span>
          </div>
        </div>
      </div>
    </div>

    <!-- Lightbox -->
    <div class="lightbox-overlay" v-if="lightboxPhoto" @click="lightboxPhoto = null">
      <span class="lightbox-close">&times;</span>
      <img :src="lightboxPhoto.image_url" :alt="lightboxPhoto.pet_name" />
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getPhotos } from '../api'
import { Camera } from 'lucide-vue-next'

const photos = ref([])
const lightboxPhoto = ref(null)

const openLightbox = (photo) => {
  lightboxPhoto.value = photo
}

onMounted(async () => {
  try {
    const res = await getPhotos()
    if (res.data.code === 200) photos.value = res.data.data
  } catch (e) {
    console.error('获取照片失败', e)
  }
})
</script>

<style lang="scss" scoped>
.gallery-section {
  background: linear-gradient(180deg, #FFFFFF, #FFF8F0);
}

.gallery-grid {
  display: grid;
  grid-template-columns: repeat(4, 1fr);
  gap: 20px;
}

.gallery-item {
  position: relative;
  border-radius: 20px;
  overflow: hidden;
  cursor: pointer;
  aspect-ratio: 1;
  box-shadow: 0 4px 15px rgba(14, 165, 233, 0.15);
  border: 4px solid white;

  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    transition: transform 0.5s ease;
  }

  .gallery-overlay {
    position: absolute;
    inset: 0;
    background: linear-gradient(to top, rgba(0,0,0,0.7) 0%, transparent 60%);
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    padding: 16px;
    opacity: 0;
    transition: opacity 0.3s ease;

    .pet-name {
      color: white;
      font-weight: 700;
      font-size: 1.1rem;
    }

    .pet-desc {
      color: rgba(255,255,255,0.8);
      font-size: 0.85rem;
    }

    .groomer {
      color: rgba(255,255,255,0.6);
      font-size: 0.8rem;
      margin-top: 4px;
    }
  }

  &:hover {
    img {
      transform: scale(1.08);
    }

    .gallery-overlay {
      opacity: 1;
    }
  }
}

@media (max-width: 768px) {
  .gallery-grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 12px;
  }
}
</style>

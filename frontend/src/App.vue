<template>
  <div id="app">
    <NavBar />
    <HeroBanner />
    <StylingSection />
    <PackageSection />
    <GallerySection />
    <ReviewSection />
    <AppointmentSection @toast="showToast" />
    <FooterSection />

    <!-- Toast 通知 -->
    <Transition name="toast">
      <div v-if="toast" class="toast" :class="toast.type">{{ toast.text }}</div>
    </Transition>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import NavBar from './components/NavBar.vue'
import HeroBanner from './components/HeroBanner.vue'
import StylingSection from './components/StylingSection.vue'
import PackageSection from './components/PackageSection.vue'
import GallerySection from './components/GallerySection.vue'
import ReviewSection from './components/ReviewSection.vue'
import AppointmentSection from './components/AppointmentSection.vue'
import FooterSection from './components/FooterSection.vue'

const toast = ref(null)
let toastTimer = null

const showToast = ({ type, text }) => {
  if (toastTimer) clearTimeout(toastTimer)
  toast.value = { type, text }
  toastTimer = setTimeout(() => { toast.value = null }, 3500)
}
</script>

<style>
.toast-enter-active { animation: slideIn 0.4s ease; }
.toast-leave-active { animation: slideIn 0.3s ease reverse; }
</style>

<template>
  <section id="appointment" class="section appointment-section">
    <div class="container">
      <h2 class="section-title"><CalendarDays class="inline-icon" :size="36" stroke-width="2.5" /> 在线预约</h2>
      <p class="section-subtitle">填写以下信息，我们会尽快与您确认预约</p>
      <div class="form-wrapper">
        <form class="appointment-form" @submit.prevent="handleSubmit">
          <div class="form-grid">
            <div class="form-group">
              <label>主人姓名 <span class="required">*</span></label>
              <input v-model="form.owner_name" type="text" placeholder="请输入您的姓名" required />
            </div>
            <div class="form-group">
              <label>联系电话 <span class="required">*</span></label>
              <input v-model="form.phone" type="tel" placeholder="请输入手机号" required />
            </div>
            <div class="form-group">
              <label>宠物种类 <span class="required">*</span></label>
              <select v-model="form.pet_type" required>
                <option value="">请选择宠物种类</option>
                <option value="狗狗">狗狗</option>
                <option value="猫咪">猫咪</option>
                <option value="其他">其他</option>
              </select>
            </div>
            <div class="form-group">
              <label>预约造型 <span class="required">*</span></label>
              <select v-model="form.styling_id" required>
                <option value="" disabled>请选择期望造型</option>
                <option v-for="s in stylings" :key="s.id" :value="String(s.id)">{{ s.name }} (¥{{ s.price }})</option>
                <option value="custom">✨ 自定义类型 (到店沟通)</option>
              </select>
            </div>
            <div class="form-group">
              <label>服务套餐 <span class="required">*</span></label>
              <select v-model="form.package_id" required>
                <option value="" disabled>请选择附加服务套餐</option>
                <option v-for="p in packages" :key="p.id" :value="String(p.id)">{{ p.name }} (¥{{ p.price }})</option>
                <option value="none">不需要额外套餐</option>
              </select>
            </div>
            <div class="form-group">
              <label>宠物名字</label>
              <input v-model="form.pet_name" type="text" placeholder="请输入宠物名字" />
            </div>
            <div class="form-group dp-container">
              <label>预约日期 <span class="required">*</span></label>
              <VueDatePicker 
                v-model="form.date" 
                :enable-time-picker="false" 
                auto-apply 
                placeholder="请选择日期" 
                :formats="{ input: 'yyyy-MM-dd', preview: 'yyyy-MM-dd' }"
                model-type="yyyy-MM-dd"
                :min-date="minDateObj"
              />
            </div>
            <div class="form-group">
              <label>预约时间 <span class="required">*</span></label>
              <select v-model="form.time" required>
                <option value="">请选择时间</option>
                <option v-for="t in timeSlots" :key="t" :value="t">{{ t }}</option>
              </select>
            </div>
          </div>
          <div class="form-group full-width">
            <label>备注信息</label>
            <textarea v-model="form.remark" placeholder="如有特殊需求请在此说明..." rows="3"></textarea>
          </div>
          <button type="submit" class="btn btn-primary submit-btn" :disabled="submitting">
            <Loader2 v-if="submitting" class="spin-icon" :size="20" />
            <CheckCircle2 v-else :size="20" />
            {{ submitting ? '提交中...' : '确认预约' }}
          </button>
        </form>
        <div class="form-side">
          <div class="info-card">
            <h4><Store class="info-icon text-primary" :size="22" stroke-width="2.5" /> 门店信息</h4>
            <ul>
              <li><MapPin :size="16" /> XX市XX区XX路100号</li>
              <li><Phone :size="16" /> 400-888-9999</li>
              <li><Clock :size="16" /> 营业时间: 9:00 - 21:00</li>
              <li><Calendar :size="16" /> 全年无休</li>
            </ul>
          </div>
          <div class="info-card">
            <h4><Info class="info-icon text-orange" :size="22" stroke-width="2.5" /> 预约须知</h4>
            <ul>
              <li><CheckCircle2 :size="16" class="text-primary" /> 请提前1天预约</li>
              <li><CheckCircle2 :size="16" class="text-primary" /> 请携带宠物健康证明</li>
              <li><CheckCircle2 :size="16" class="text-primary" /> 如需取消请提前4小时</li>
              <li><CheckCircle2 :size="16" class="text-primary" /> 首次到店享过折扣优惠</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { createAppointment, getStylings, getPackages } from '../api'
import { CalendarDays, Store, Info, MapPin, Phone, Clock, Calendar, CheckCircle2, Loader2 } from 'lucide-vue-next'
import { VueDatePicker } from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'

const emit = defineEmits(['toast'])

const submitting = ref(false)

const form = ref({
  owner_name: '',
  phone: '',
  pet_type: '',
  pet_name: '',
  styling_id: '',
  package_id: '',
  date: '',
  time: '',
  remark: ''
})

const stylings = ref([])
const packages = ref([])

onMounted(async () => {
  try {
    const res = await Promise.all([getStylings(), getPackages()])
    if (res[0].data.code === 200) stylings.value = res[0].data.data
    if (res[1].data.code === 200) packages.value = res[1].data.data
  } catch (e) {
    console.error('获取造型或套餐列表失败', e)
  }
})

const minDateObj = computed(() => {
  const d = new Date()
  d.setDate(d.getDate() + 1)
  return d
})

const timeSlots = [
  '09:00', '09:30', '10:00', '10:30', '11:00', '11:30',
  '13:00', '13:30', '14:00', '14:30', '15:00', '15:30',
  '16:00', '16:30', '17:00', '17:30', '18:00', '18:30',
  '19:00', '19:30', '20:00'
]

const handleSubmit = async () => {
  // 手机号校验
  if (!/^1\d{10}$/.test(form.value.phone)) {
    emit('toast', { type: 'error', text: '请输入正确的手机号码' })
    return
  }

  submitting.value = true
  try {
    const payload = {
      ...form.value,
      styling_id: String(form.value.styling_id),
      package_id: String(form.value.package_id)
    }
    const res = await createAppointment(payload)
    if (res.data.code === 200) {
      emit('toast', { type: 'success', text: '🎉 预约成功！我们会尽快与您确认' })
      // 重置表单
      form.value = { owner_name: '', phone: '', pet_type: '', pet_name: '', styling_id: '', package_id: '', date: '', time: '', remark: '' }
    } else {
      emit('toast', { type: 'error', text: res.data.message || '预约失败，请重试' })
    }
  } catch (e) {
    emit('toast', { type: 'error', text: '网络异常，请稍后重试' })
  } finally {
    submitting.value = false
  }
}
</script>

<style lang="scss" scoped>
.appointment-section {
  background: linear-gradient(180deg, #FFFFFF, #F0F9FF);
}

.form-wrapper {
  display: grid;
  grid-template-columns: 1fr 320px;
  gap: 32px;
  align-items: start;
}

.appointment-form {
  background: white;
  padding: 40px;
  border-radius: 24px;
  box-shadow: 12px 12px 24px rgba(14, 165, 233, 0.1), -12px -12px 24px rgba(255, 255, 255, 0.9);
  border: 4px solid white;
}

.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
  margin-bottom: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;

  &.full-width {
    margin-bottom: 24px;
  }

  label {
    font-weight: 700;
    font-size: 0.9rem;
    color: #4A3728;

    .required {
      color: #FF6B6B;
    }
  }

  input, select, textarea {
    padding: 14px 18px;
    border: 3px solid #F1F5F9;
    border-radius: 16px;
    font-family: 'Nunito', sans-serif;
    font-size: 1rem;
    background: #F8FAFC;
    color: #0C4A6E;
    transition: all 0.2s cubic-bezier(0.34, 1.56, 0.64, 1);
    outline: none;
    box-shadow: inset 2px 2px 5px rgba(14, 165, 233, 0.05);

    &:focus {
      border-color: #38BDF8;
      background: white;
      box-shadow: 0 0 0 4px rgba(56, 189, 248, 0.2), inset 1px 1px 2px transparent;
      transform: translateY(-2px);
    }

    &::placeholder {
      color: #94A3B8;
    }
  }

  textarea {
    resize: vertical;
  }
}

/* 覆盖 vue-datepicker 默认样式，契合黏土风 */
:deep(.dp__theme_light) {
  --dp-background-color: #F8FAFC;
  --dp-text-color: #0C4A6E;
  --dp-border-color: #F1F5F9;
  --dp-border-color-hover: #38BDF8;
  --dp-primary-color: #0EA5E9;
  --dp-border-radius: 16px;
}

:deep(.dp__input) {
  padding: 14px 18px 14px 40px; /* 增加左侧 padding，避免文字被原生图标遮挡 */
  border: 3px solid #F1F5F9;
  border-radius: 16px;
  font-family: 'Nunito', sans-serif;
  font-size: 1rem;
  background: #F8FAFC;
  color: #0C4A6E;
  transition: all 0.2s cubic-bezier(0.34, 1.56, 0.64, 1);
  box-shadow: inset 2px 2px 5px rgba(14, 165, 233, 0.05);

  &:focus {
    border-color: #38BDF8;
    background: white;
    box-shadow: 0 0 0 4px rgba(56, 189, 248, 0.2), inset 1px 1px 2px transparent;
    transform: translateY(-2px);
  }
  
  &::placeholder {
    color: #94A3B8;
  }
}

.submit-btn {
  width: 100%;
  padding: 14px;
  font-size: 1.05rem;
  justify-content: center;

  &:disabled {
    opacity: 0.6;
    cursor: not-allowed;
    transform: none !important;
  }
}

.form-side {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.info-card {
  background: white;
  padding: 28px;
  border-radius: 24px;
  box-shadow: 8px 8px 16px rgba(14, 165, 233, 0.1), -8px -8px 16px rgba(255, 255, 255, 0.8);
  border: 3px solid #F0F9FF;

  h4 {
    font-size: 1.2rem;
    margin-bottom: 16px;
    color: #0C4A6E;
    display: flex;
    align-items: center;
    gap: 8px;
  }

  ul {
    list-style: none;

    li {
      padding: 8px 0;
      font-size: 0.95rem;
      color: #475569;
      display: flex;
      align-items: center;
      gap: 10px;
    }
  }
}

.spin-icon {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

@media (max-width: 768px) {
  .form-wrapper {
    grid-template-columns: 1fr;
  }

  .form-grid {
    grid-template-columns: 1fr;
  }

  .appointment-form {
    padding: 24px;
  }
}
</style>

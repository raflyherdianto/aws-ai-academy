<template>
  <div class="w-full max-w-md mx-auto">
    <div class="glass-container p-8 relative overflow-hidden">
      <div class="absolute -top-10 -right-10 w-32 h-32 bg-aws-orange/10 rounded-full blur-xl"></div>
      <div class="absolute -bottom-10 -left-10 w-32 h-32 bg-dicoding-blue/15 rounded-full blur-xl"></div>
      <div class="text-center mb-8 relative z-10">
        <h2 class="text-2xl font-bold text-white mb-2">Mulai Berkenalan!</h2>
        <p class="text-slate-400 text-sm">
          AWS AI Academy Cohort Networking. Masukkan email terdaftar Anda untuk membuat atau melihat kartu karakter Anda.
        </p>
      </div>
      <form @submit.prevent="handleVerify" class="space-y-6 relative z-10">
        <div>
          <label for="email" class="block text-sm font-medium text-slate-300 mb-2">Email Peserta / Fasil</label>
          <input 
            type="email" 
            id="email" 
            v-model="email" 
            required 
            placeholder="nama@email.com" 
            class="w-full px-4 py-3 bg-white/5 border border-white/10 rounded-xl text-white placeholder-slate-500 focus:outline-none focus:border-dicoding-blue focus:ring-2 focus:ring-dicoding-blue/20 transition-all text-sm"
            :disabled="loading"
          />
        </div>
        <button 
          type="submit" 
          class="w-full py-3 bg-dicoding-blue hover:bg-dicoding-blue/90 text-white font-medium rounded-xl transition-all shadow-lg shadow-dicoding-blue/20 flex items-center justify-center gap-2 text-sm cursor-pointer"
          :disabled="loading"
        >
          <span v-if="loading" class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></span>
          <span>{{ loading ? 'Memverifikasi...' : 'Mulai Sekarang' }}</span>
        </button>
      </form>
      <div v-if="error" class="mt-4 p-3 bg-rose-500/10 border border-rose-500/20 text-rose-300 text-xs rounded-xl text-center relative z-10">
        {{ error }}
      </div>
      <div class="mt-8 pt-6 border-t border-white/5 text-center relative z-10">
        <p class="text-slate-500 text-xs mb-3">Hanya ingin bermain game tebak-tebakan?</p>
        <router-link 
          to="/game" 
          class="inline-flex items-center gap-1.5 text-xs text-aws-orange hover:underline font-medium"
        >
          Mainkan Game "Guess Who" 
          <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="5" y1="12" x2="19" y2="12"></line><polyline points="12 5 19 12 12 19"></polyline></svg>
        </router-link>
      </div>
    </div>
  </div>
</template>
<script>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
export default {
  name: 'HomeView',
  setup() {
    const email = ref('')
    const loading = ref(false)
    const error = ref('')
    const router = useRouter()
    const handleVerify = async () => {
      loading.value = true
      error.value = ''
      try {
        const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
        const response = await fetch(`${apiUrl}/api/auth?email=${encodeURIComponent(email.value.trim())}`)
        const data = await response.json()
        if (!response.ok) {
          error.value = data.message || 'Email tidak ditemukan di daftar peserta.'
          return
        }
        localStorage.setItem('user_email', email.value.trim())
        if (data.token) {
          localStorage.setItem('jwt_token', data.token)
        }
        const slugify = (text) => {
          if (!text) return ''
          return text.toString().toLowerCase()
            .replace(/\s+/g, '-')
            .replace(/[^\w\-]+/g, '')
            .replace(/\-\-+/g, '-')
            .replace(/^-+/, '')
            .replace(/-+$/, '')
        }
        if (data.is_registered) {
          const nameSlug = slugify(data.participant.nama_lengkap || '')
          router.push(`/share/${data.participant.id}/${nameSlug}`)
        } else {
          router.push('/register')
        }
      } catch (err) {
        error.value = 'Gagal terhubung ke server backend. Pastikan server Go berjalan.'
        console.error(err)
      } finally {
        loading.value = false
      }
    }
    return {
      email,
      loading,
      error,
      handleVerify
    }
  }
}
</script>

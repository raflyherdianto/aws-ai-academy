<template>
  <div class="w-full max-w-lg mx-auto flex flex-col items-center gap-6">
    <div v-if="loading" class="flex flex-col items-center gap-2 py-12">
      <span class="w-8 h-8 border-4 border-aws-orange border-t-transparent rounded-full animate-spin"></span>
      <span class="text-sm text-slate-400">Memuat kartu karakter...</span>
    </div>
    <div v-else-if="error" class="glass-container p-8 text-center w-full">
      <p class="text-rose-400 font-semibold mb-4">{{ error }}</p>
      <router-link to="/" class="px-4 py-2 bg-white/5 hover:bg-white/10 text-white rounded-xl text-sm transition-all">
        Kembali ke Beranda
      </router-link>
    </div>
    <div v-else class="w-full flex flex-col items-center gap-6">
      <div class="text-center">
        <h2 class="text-xl font-bold text-white mb-1">Kartu Karakter Tercipta! 🎉</h2>
        <p class="text-xs text-slate-400">Bagikan kartu Anda ke WhatsApp Group untuk berjejaring.</p>
      </div>
      <div 
        ref="rpgCard"
        class="w-[340px] h-[480px] sm:w-[420px] sm:h-[580px] p-5 sm:p-7 glass-card relative overflow-hidden flex flex-col justify-between shadow-2xl shadow-black/60 rounded-2xl border border-white/10"
        style="background: linear-gradient(135deg, rgba(15, 23, 42, 0.9) 0%, rgba(30, 41, 59, 0.8) 100%);"
      >
        <div class="absolute -top-24 -left-24 w-48 h-48 bg-aws-orange/10 rounded-full blur-2xl"></div>
        <div class="absolute -bottom-24 -right-24 w-48 h-48 bg-dicoding-blue/15 rounded-full blur-2xl"></div>
        <div class="flex justify-between items-center border-b border-white/10 pb-3 relative z-10 w-full">
          <div class="flex items-center gap-3">
            <img 
              :src="avatarDataUrl || `https://api.dicebear.com/7.x/pixel-art/svg?seed=${encodeURIComponent(participant.email || 'adventurer')}`"
              class="w-10 h-10 sm:w-12 sm:h-12 rounded-lg border border-white/10 bg-slate-950/40 p-0.5"
              alt="Avatar"
            />
            <div class="flex flex-col justify-center">
              <span class="text-[8px] sm:text-[9px] font-bold text-aws-orange tracking-wider uppercase font-mono leading-none mb-1">AWS AI ACADEMY</span>
              <h3 class="text-[13px] sm:text-[15px] font-bold text-white font-mono tracking-tight leading-none capitalize">
                {{ participant.nama_panggilan || participant.nama_lengkap }}
              </h3>
            </div>
          </div>
          <span class="px-2 py-0.5 sm:px-2.5 sm:py-1 bg-white/5 border border-white/10 rounded text-[8px] sm:text-[9px] font-bold text-slate-300 uppercase font-mono">
            {{ participant.rpg_class || 'Adventurer' }}
          </span>
        </div>
        <div class="my-3 sm:my-4 space-y-2 sm:space-y-3 relative z-10 flex-grow flex flex-col justify-start">
          <div class="text-center py-2 px-3 sm:py-2.5 sm:px-4 bg-white/5 rounded-xl border border-white/5">
            <p class="text-slate-300 text-xs sm:text-[13px] italic line-clamp-3 leading-relaxed">
              "{{ participant.bio || 'Hello World!' }}"
            </p>
          </div>
          <div class="grid grid-cols-2 gap-2 sm:gap-2.5 text-[10px] sm:text-[11px]">
            <div class="bg-white/5 p-2 sm:p-2.5 rounded-lg border border-white/5 flex flex-col justify-between">
              <span class="text-aws-orange text-[8px] sm:text-[9px] block uppercase font-mono font-bold tracking-wider">PEKERJAAN</span>
              <span class="text-slate-200 font-semibold truncate block">{{ participant.pekerjaan || '-' }}</span>
            </div>
            <div class="bg-white/5 p-2 sm:p-2.5 rounded-lg border border-white/5 flex flex-col justify-between">
              <span class="text-aws-orange text-[8px] sm:text-[9px] block uppercase font-mono font-bold tracking-wider">DOMISILI</span>
              <span class="text-slate-200 font-semibold truncate block">
                {{ participant.kota ? `${participant.kota}, ${participant.provinsi}` : '-' }}
              </span>
            </div>
          </div>
          <div v-if="participant.motivasi" class="p-2 sm:p-2.5 bg-white/5 border border-white/5 rounded-lg text-left">
            <span class="text-aws-orange block font-bold uppercase tracking-wider text-[7px] sm:text-[8px] mb-0.5">MOTIVASI</span>
            <p class="text-slate-300 text-[9.5px] sm:text-[10.5px] leading-relaxed line-clamp-2">
              {{ participant.motivasi }}
            </p>
          </div>
          <div v-if="participant.linkedin" class="p-2 sm:p-2.5 bg-white/5 border border-white/5 rounded-lg text-left flex items-start gap-2">
            <svg class="w-3.5 h-3.5 sm:w-4 sm:h-4 text-blue-400 shrink-0 mt-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"></path><rect x="2" y="9" width="4" height="12"></rect><circle cx="4" cy="4" r="2"></circle></svg>
            <span class="text-slate-300 text-[9.5px] sm:text-[11px] break-all font-mono leading-tight">linkedin.com/in/{{ cleanLinkedInUrl(participant.linkedin) }}</span>
          </div>
        </div>
        <div class="space-y-2 border-t border-white/10 pt-3 sm:pt-4 relative z-10">
          <div class="space-y-1.5 sm:space-y-2">
            <div v-for="(val, stat) in stats" :key="stat" class="space-y-0.5">
              <div class="flex justify-between text-[9px] sm:text-[10.5px] font-mono font-semibold">
                <span class="text-slate-400 capitalize">{{ stat === 'bugs' ? 'Bug Resistance' : stat + ' Power' }}</span>
                <span class="text-white">{{ val }}</span>
              </div>
              <div class="w-full h-1.5 sm:h-2 bg-slate-900 rounded-full overflow-hidden border border-white/5">
                <div 
                  class="h-full rounded-full transition-all duration-500" 
                  :class="stat === 'coding' || stat === 'architecture' ? 'bg-dicoding-blue' : 'bg-aws-orange'"
                  :style="{ width: `${val}%` }"
                ></div>
              </div>
            </div>
          </div>
        </div>
        <div class="flex justify-between items-center text-[8px] sm:text-[9.5px] font-mono text-slate-500 border-t border-white/5 pt-2 mt-2 relative z-10">
          <span>ID: FASIL-RAFLY-2026</span>
          <span>AWS x DICODING</span>
        </div>
      </div>
      <div class="w-full glass-container p-6 space-y-4">
        <div class="grid gap-3" :class="participant.linkedin ? 'grid-cols-3' : 'grid-cols-2'">
          <button 
            @click="copyLink"
            class="py-2.5 px-4 bg-white/5 hover:bg-white/10 text-white rounded-xl border border-white/10 text-xs font-semibold flex items-center justify-center gap-2 cursor-pointer transition-all"
          >
            <span>🔗</span> {{ copied ? 'Tersalin!' : 'Salin Tautan' }}
          </button>
          <a 
            :href="whatsappShareUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="py-2.5 px-4 bg-emerald-600 hover:bg-emerald-500 text-white rounded-xl text-xs font-semibold flex items-center justify-center gap-2 cursor-pointer transition-all text-center"
          >
            <span>💬</span> Share WA
          </a>
          <a 
            v-if="participant.linkedin"
            :href="participant.linkedin"
            target="_blank"
            rel="noopener noreferrer"
            class="py-2.5 px-4 bg-blue-700 hover:bg-blue-600 text-white rounded-xl text-xs font-semibold flex items-center justify-center gap-2 cursor-pointer transition-all text-center"
          >
            <span>💼</span> LinkedIn
          </a>
        </div>
        <div class="flex flex-col gap-2 pt-2 border-t border-white/5">
          <button 
            @click="downloadCard"
            :disabled="downloading"
            class="w-full py-2.5 bg-aws-orange hover:bg-aws-orange/90 text-slate-950 font-bold rounded-xl text-xs flex items-center justify-center gap-2 cursor-pointer transition-all text-center"
          >
            <span>📥</span> {{ downloading ? 'Rendering HD...' : 'Unduh Kartu (HD)' }}
          </button>
        </div>
        <div class="flex justify-between text-xs pt-4 border-t border-white/5">
          <router-link to="/" class="text-slate-400 hover:text-white transition-colors">
            ← Buat Kartu Baru
          </router-link>
          <router-link to="/game" class="text-aws-orange hover:underline font-semibold">
            Main Game Tebak Profil →
          </router-link>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { ref, onMounted, computed } from 'vue'
import { useRoute } from 'vue-router'
import html2canvas from 'html2canvas-pro'
export default {
  name: 'ShareView',
  setup() {
    const route = useRoute()
    const participant = ref({})
    const stats = ref({})
    const loading = ref(true)
    const error = ref('')
    const copied = ref(false)
    const avatarDataUrl = ref('')
    const fetchProfile = async () => {
      loading.value = true
      error.value = ''
      try {
        const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
        const response = await fetch(`${apiUrl}/api/profile/${route.params.id}`)
        if (!response.ok) {
          error.value = 'Kartu karakter tidak ditemukan.'
          return
        }
        const data = await response.json()
        participant.value = data
        if (data.rpg_stats) {
          stats.value = JSON.parse(data.rpg_stats)
        }
        if (data.email) {
          try {
            const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
            const pngRes = await fetch(`${apiUrl}/api/avatar?seed=${encodeURIComponent(data.email)}`)
            if (pngRes.ok) {
              const blob = await pngRes.blob()
              const reader = new FileReader()
              reader.onloadend = () => {
                avatarDataUrl.value = reader.result 
              }
              reader.readAsDataURL(blob)
            }
          } catch (avatarErr) {
            console.error('Gagal mengambil avatar PNG:', avatarErr)
          }
        }
      } catch (err) {
        error.value = 'Gagal terhubung ke server backend.'
        console.error(err)
      } finally {
        loading.value = false
      }
    }
    onMounted(fetchProfile)
    const slugify = (text) => {
      if (!text) return ''
      return text.toString().toLowerCase()
        .replace(/\s+/g, '-')
        .replace(/[^\w\-]+/g, '')
        .replace(/\-\-+/g, '-')
        .replace(/^-+/, '')
        .replace(/-+$/, '')
    }
    const shareUrl = computed(() => {
      const id = participant.value.id || route.params.id
      const nameSlug = slugify(participant.value.nama_lengkap || '')
      return `${window.location.origin}/share/${id}/${nameSlug}`
    })
    const capitalizeName = (str) => {
      if (!str) return ''
      return str.split(' ').map(w => w.charAt(0).toUpperCase() + w.slice(1).toLowerCase()).join(' ')
    }
    const whatsappShareUrl = computed(() => {
      const name = capitalizeName(participant.value.nama_panggilan || participant.value.nama_lengkap || '')
      const text = `Halo! Yuk intip kartu karakter developer ${name} di AWS AI Academy: ${shareUrl.value}`
      return `https://api.whatsapp.com/send?text=${encodeURIComponent(text)}`
    })
    const downloadImageUrl = computed(() => {
      if (participant.value.image_path) {
        const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
        return `${apiUrl}${participant.value.image_path}`
      }
      return '#'
    })
    const cleanLinkedInUrl = (url) => {
      if (!url) return ''
      return url.replace(/https?:\/\/(www\.)?linkedin\.com\/in\/?/, '').replace(/\/$/, '')
    }
    const rpgCard = ref(null)
    const downloading = ref(false)
    const downloadCard = async () => {
      if (!rpgCard.value) return
      downloading.value = true
      try {
        const canvas = await html2canvas(rpgCard.value, {
          scale: 3, 
          useCORS: true,
          backgroundColor: null
        })
        const image = canvas.toDataURL('image/png')
        const link = document.createElement('a')
        const nickname = participant.value.nama_panggilan || 'developer'
        link.download = `card_${nickname}.png`
        link.href = image
        link.click()
      } catch (err) {
        console.error('Gagal mengunduh kartu:', err)
        alert('Gagal mengunduh kartu: ' + err.message)
      } finally {
        downloading.value = false
      }
    }
    const copyLink = async () => {
      try {
        const id = participant.value.id || route.params.id
        const nameSlug = slugify(participant.value.nama_lengkap || '')
        const frontendShareUrl = `${window.location.origin}/share/${id}/${nameSlug}`
        await navigator.clipboard.writeText(frontendShareUrl)
        copied.value = true
        setTimeout(() => { copied.value = false }, 2000)
      } catch (err) {
        console.error('Gagal menyalin link', err)
      }
    }
    return {
      participant,
      stats,
      loading,
      error,
      copied,
      whatsappShareUrl,
      downloadImageUrl,
      copyLink,
      rpgCard,
      downloading,
      downloadCard,
      avatarDataUrl,
      cleanLinkedInUrl
    }
  }
}
</script>

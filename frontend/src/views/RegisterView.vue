<template>
  <div class="w-full max-w-5xl mx-auto grid grid-cols-1 lg:grid-cols-2 gap-8 items-start">
    <div class="glass-container p-6 md:p-8 relative overflow-hidden">
      <div class="absolute -top-10 -left-10 w-24 h-24 bg-aws-orange/10 rounded-full blur-xl"></div>
      <h2 class="text-xl font-bold text-white mb-6 flex items-center gap-2">
        <span class="text-aws-orange">⚡</span> Lengkapi Profil Coder Anda
      </h2>
      <form @submit.prevent="handleSubmit" class="space-y-5">
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label class="block text-xs font-semibold text-slate-300 mb-1.5">Nama Lengkap</label>
            <input 
              type="text" 
              :value="fullName" 
              disabled 
              class="w-full px-3 py-2 bg-white/5 border border-white/5 rounded-lg text-slate-400 text-sm cursor-not-allowed"
            />
          </div>
          <div>
            <label class="block text-xs font-semibold text-slate-300 mb-1.5">Nama Panggilan</label>
            <input 
              type="text" 
              v-model="form.nama_panggilan" 
              required 
              placeholder="e.g. Budi"
              class="w-full px-3 py-2 bg-white/5 border border-white/10 rounded-lg text-white focus:outline-none focus:border-dicoding-blue text-sm"
            />
          </div>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label class="block text-xs font-semibold text-slate-300 mb-1.5">Gender</label>
            <select 
              v-model="form.gender" 
              required 
              class="w-full px-3 py-2 bg-slate-900 border border-white/10 rounded-lg text-white focus:outline-none focus:border-dicoding-blue text-sm"
            >
              <option value="" disabled>Pilih...</option>
              <option value="Laki-laki">Laki-laki</option>
              <option value="Perempuan">Perempuan</option>
            </select>
          </div>
          <div>
            <label class="block text-xs font-semibold text-slate-300 mb-1.5">Pekerjaan Saat Ini</label>
            <input 
              type="text" 
              v-model="form.pekerjaan" 
              required 
              placeholder="e.g. Mahasiswa, Backend Dev"
              class="w-full px-3 py-2 bg-white/5 border border-white/10 rounded-lg text-white focus:outline-none focus:border-dicoding-blue text-sm"
            />
          </div>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div class="relative" ref="provinceDropdownRef">
            <label class="block text-xs font-semibold text-slate-300 mb-1.5">Provinsi</label>
            <div
              @click="toggleProvinceDropdown"
              class="w-full px-3 py-2 bg-white/5 border border-white/10 rounded-lg text-white focus:outline-none focus:border-dicoding-blue text-sm cursor-pointer flex justify-between items-center"
              :class="provinceOpen ? 'border-dicoding-blue' : ''"
            >
              <span :class="form.provinsi ? 'text-white' : 'text-slate-500'">{{ form.provinsi || 'Pilih Provinsi...' }}</span>
              <svg class="w-4 h-4 text-slate-400 transition-transform" :class="provinceOpen ? 'rotate-180' : ''" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/></svg>
            </div>
            <div v-if="provinceOpen" class="absolute z-50 top-full mt-1 left-0 right-0 bg-slate-900 border border-white/10 rounded-lg shadow-xl overflow-hidden">
              <div class="p-2">
                <input
                  v-model="provinceSearch"
                  @click.stop
                  placeholder="Cari provinsi..."
                  class="w-full px-2 py-1.5 bg-white/5 border border-white/10 rounded text-white text-xs focus:outline-none focus:border-dicoding-blue"
                  autofocus
                />
              </div>
              <ul class="max-h-48 overflow-y-auto">
                <li
                  v-for="p in filteredProvinces"
                  :key="p.code"
                  @click="selectProvince(p)"
                  class="px-3 py-2 text-xs text-slate-300 hover:bg-white/10 cursor-pointer transition-colors"
                  :class="selectedProvince?.code === p.code ? 'bg-aws-orange/20 text-aws-orange' : ''"
                >
                  {{ p.name }}
                </li>
                <li v-if="filteredProvinces.length === 0" class="px-3 py-2 text-xs text-slate-500">Tidak ditemukan</li>
              </ul>
            </div>
          </div>
          <div class="relative" ref="regencyDropdownRef">
            <label class="block text-xs font-semibold text-slate-300 mb-1.5">Kota / Kabupaten</label>
            <div
              @click="toggleRegencyDropdown"
              class="w-full px-3 py-2 bg-white/5 border border-white/10 rounded-lg text-sm cursor-pointer flex justify-between items-center"
              :class="[!selectedProvince ? 'opacity-50 cursor-not-allowed text-slate-500' : regencyOpen ? 'border-dicoding-blue text-white' : 'text-white', regencyOpen ? 'border border-dicoding-blue' : 'border border-white/10']"
            >
              <span :class="form.kota ? 'text-white' : 'text-slate-500'">{{ form.kota || (selectedProvince ? 'Pilih Kota...' : 'Pilih provinsi dulu') }}</span>
              <svg class="w-4 h-4 text-slate-400 transition-transform" :class="regencyOpen ? 'rotate-180' : ''" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/></svg>
            </div>
            <div v-if="regencyOpen && selectedProvince" class="absolute z-50 top-full mt-1 left-0 right-0 bg-slate-900 border border-white/10 rounded-lg shadow-xl overflow-hidden">
              <div class="p-2">
                <input
                  v-model="regencySearch"
                  @click.stop
                  placeholder="Cari kota/kabupaten..."
                  class="w-full px-2 py-1.5 bg-white/5 border border-white/10 rounded text-white text-xs focus:outline-none focus:border-dicoding-blue"
                  autofocus
                />
              </div>
              <ul class="max-h-48 overflow-y-auto">
                <li
                  v-for="r in filteredRegencies"
                  :key="r.code"
                  @click="selectRegency(r)"
                  class="px-3 py-2 text-xs text-slate-300 hover:bg-white/10 cursor-pointer transition-colors"
                  :class="form.kota === r.name ? 'bg-aws-orange/20 text-aws-orange' : ''"
                >
                  {{ r.name }}
                </li>
                <li v-if="filteredRegencies.length === 0" class="px-3 py-2 text-xs text-slate-500">Tidak ditemukan</li>
              </ul>
            </div>
          </div>
        </div>
        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
          <div>
            <label class="block text-xs font-semibold text-slate-300 mb-1.5">LinkedIn Profile (Opsional)</label>
            <input 
              type="url" 
              v-model="form.linkedin" 
              placeholder="https://linkedin.com/in/username"
              class="w-full px-3 py-2 bg-white/5 border border-white/10 rounded-lg text-white focus:outline-none focus:border-dicoding-blue text-sm"
            />
          </div>
          <div>
            <label class="block text-xs font-semibold text-slate-300 mb-1.5">Background IT / Non-IT?</label>
            <select 
              v-model="form.background_it" 
              required 
              class="w-full px-3 py-2 bg-slate-900 border border-white/10 rounded-lg text-white focus:outline-none focus:border-dicoding-blue text-sm"
            >
              <option value="" disabled>Pilih...</option>
              <option value="IT">IT</option>
              <option value="Non-IT">Non-IT</option>
            </select>
          </div>
        </div>
        <div>
          <label class="block text-xs font-semibold text-slate-300 mb-1.5">Coder Class</label>
          <div class="grid grid-cols-2 gap-2">
            <button 
              v-for="cls in classes" 
              :key="cls.name"
              type="button"
              @click="form.rpg_class = cls.name"
              class="py-2.5 px-3 rounded-lg border text-xs font-semibold transition-all cursor-pointer"
              :class="form.rpg_class === cls.name ? 'border-aws-orange bg-aws-orange/15 text-aws-orange' : 'border-white/10 bg-white/5 text-slate-300 hover:bg-white/10'"
            >
              {{ cls.name }}
            </button>
          </div>
        </div>
        <div class="p-4 bg-white/5 border border-white/10 rounded-xl">
          <div class="flex justify-between items-center mb-3">
            <div>
              <span class="text-xs font-semibold text-slate-200">Coder Stats</span>
              <p class="text-[10px] text-slate-500 mt-0.5">Dihitung berdasarkan profil Anda</p>
            </div>
            <button 
              type="button" 
              @click="calculateAdaptiveStats"
              class="px-2.5 py-1 bg-aws-orange hover:bg-aws-orange/90 text-slate-950 font-bold rounded text-[10px] uppercase tracking-wide cursor-pointer transition-all"
            >
              ⚡ Kalkulasi
            </button>
          </div>
          <div class="grid grid-cols-2 gap-3 text-xs">
            <div class="flex justify-between">
              <span class="text-slate-400">Coding Power:</span>
              <span class="font-bold text-white font-mono">{{ form.rpg_stats.coding }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-slate-400">Bug Resistance:</span>
              <span class="font-bold text-white font-mono">{{ form.rpg_stats.bugs }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-slate-400">Architecture:</span>
              <span class="font-bold text-white font-mono">{{ form.rpg_stats.architecture }}</span>
            </div>
            <div class="flex justify-between">
              <span class="text-slate-400">Cloud Affinity:</span>
              <span class="font-bold text-white font-mono">{{ form.rpg_stats.cloud }}</span>
            </div>
          </div>
        </div>
        <div>
          <label class="block text-xs font-semibold text-slate-300 mb-1.5">Bio Perkenalan Singkat</label>
          <textarea 
            v-model="form.bio" 
            required 
            rows="2"
            maxlength="150"
            placeholder="Tulis bio unik perkenalan Anda (maksimal 150 karakter)..."
            class="w-full px-3 py-2 bg-white/5 border border-white/10 rounded-lg text-white focus:outline-none focus:border-dicoding-blue text-sm resize-none"
          ></textarea>
        </div>
        <div>
          <label class="block text-xs font-semibold text-slate-300 mb-1.5">Motivasi Mengikuti AWS AI Academy</label>
          <textarea 
            v-model="form.motivasi" 
            required 
            rows="2"
            maxlength="150"
            placeholder="Tulis motivasi Anda mengikuti program ini..."
            class="w-full px-3 py-2 bg-white/5 border border-white/10 rounded-lg text-white focus:outline-none focus:border-dicoding-blue text-sm resize-none"
          ></textarea>
        </div>
        <button 
          type="submit" 
          class="w-full py-3 bg-aws-orange hover:bg-aws-orange/90 text-slate-950 font-bold rounded-xl transition-all shadow-lg shadow-aws-orange/20 flex items-center justify-center gap-2 cursor-pointer text-sm"
          :disabled="submitting"
        >
          <span v-if="submitting" class="w-4 h-4 border-2 border-slate-950 border-t-transparent rounded-full animate-spin"></span>
          <span>{{ submitting ? 'Memproses Kartu...' : 'Simpan & Dapatkan Kartu Developer' }}</span>
        </button>
      </form>
      <p v-if="error" class="mt-3 text-rose-400 text-xs text-center">{{ error }}</p>
    </div>
    <div class="flex flex-col items-center gap-4">
      <span class="text-xs font-bold uppercase tracking-wider text-slate-400">Pratinjau Kartu Karakter</span>
      <div 
        ref="rpgCard"
        class="w-[340px] h-[480px] sm:w-[420px] sm:h-[580px] p-5 sm:p-7 glass-card relative overflow-hidden flex flex-col justify-between shadow-2xl shadow-black/60 rounded-2xl border border-white/10 select-none"
        style="background: linear-gradient(135deg, rgba(15, 23, 42, 0.9) 0%, rgba(30, 41, 59, 0.8) 100%);"
      >
        <div class="absolute -top-24 -left-24 w-48 h-48 bg-aws-orange/10 rounded-full blur-2xl"></div>
        <div class="absolute -bottom-24 -right-24 w-48 h-48 bg-dicoding-blue/15 rounded-full blur-2xl"></div>
        <div class="flex justify-between items-center border-b border-white/10 pb-3 relative z-10 w-full">
          <div class="flex items-center gap-3">
            <img 
              :src="avatarDataUrl || `https://api.dicebear.com/7.x/pixel-art/svg?seed=${encodeURIComponent(form.email || 'adventurer')}`"
              class="w-10 h-10 sm:w-12 sm:h-12 rounded-lg border border-white/10 bg-slate-950/40 p-0.5"
              alt="Avatar"
            />
            <div class="flex flex-col justify-center">
              <span class="text-[8px] sm:text-[9px] font-bold text-aws-orange tracking-wider uppercase font-mono leading-none mb-1">AWS AI ACADEMY</span>
              <h3 class="text-[13px] sm:text-[15px] font-bold text-white font-mono tracking-tight leading-none capitalize">
                {{ form.nama_panggilan || fullName || 'Adventurer' }}
              </h3>
            </div>
          </div>
          <span class="px-2 py-0.5 sm:px-2.5 sm:py-1 bg-white/5 border border-white/10 rounded text-[8px] sm:text-[9px] font-bold text-slate-300 uppercase font-mono">
            {{ form.rpg_class || 'No Class' }}
          </span>
        </div>
        <div class="my-3 sm:my-4 space-y-2 sm:space-y-3 relative z-10 flex-grow flex flex-col justify-start">
          <div class="text-center py-2 px-3 sm:py-2.5 sm:px-4 bg-white/5 rounded-xl border border-white/5">
            <p class="text-slate-300 text-xs sm:text-[13px] italic line-clamp-3 leading-relaxed">
              "{{ form.bio || 'Tulis bio perkenalan Anda di form sebelah kiri...' }}"
            </p>
          </div>
          <div class="grid grid-cols-2 gap-2 sm:gap-2.5 text-[10px] sm:text-[11px]">
            <div class="bg-white/5 p-2 sm:p-2.5 rounded-lg border border-white/5 flex flex-col justify-between">
              <span class="text-aws-orange text-[8px] sm:text-[9px] block uppercase font-mono font-bold tracking-wider">PEKERJAAN</span>
              <span class="text-slate-200 font-semibold truncate block">{{ form.pekerjaan || '-' }}</span>
            </div>
            <div class="bg-white/5 p-2 sm:p-2.5 rounded-lg border border-white/5 flex flex-col justify-between">
              <span class="text-aws-orange text-[8px] sm:text-[9px] block uppercase font-mono font-bold tracking-wider">DOMISILI</span>
              <span class="text-slate-200 font-semibold truncate block">
                {{ form.kota ? `${form.kota}, ${form.provinsi}` : '-' }}
              </span>
            </div>
          </div>
          <div class="mt-1 p-2 sm:p-2.5 bg-white/5 border border-white/5 rounded-lg text-left">
            <span class="text-aws-orange block font-bold uppercase tracking-wider text-[7px] sm:text-[8px] mb-0.5">MOTIVASI</span>
            <p class="text-slate-300 text-[9.5px] sm:text-[10.5px] leading-relaxed line-clamp-2">
              {{ form.motivasi || 'Tulis motivasi Anda di form sebelah kiri...' }}
            </p>
          </div>
          <div v-if="form.linkedin" class="p-2 sm:p-2.5 bg-white/5 border border-white/5 rounded-lg text-left flex items-start gap-2">
            <svg class="w-3.5 h-3.5 sm:w-4 sm:h-4 text-blue-400 shrink-0 mt-0.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2 2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6z"></path><rect x="2" y="9" width="4" height="12"></rect><circle cx="4" cy="4" r="2"></circle></svg>
            <span class="text-slate-300 text-[9.5px] sm:text-[11px] break-all font-mono leading-tight">linkedin.com/in/{{ cleanLinkedInUrl(form.linkedin) }}</span>
          </div>
        </div>
        <div class="space-y-2 border-t border-white/10 pt-3 sm:pt-4 relative z-10">
          <div class="space-y-1.5 sm:space-y-2">
            <div v-for="(val, stat) in form.rpg_stats" :key="stat" class="space-y-0.5">
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
    </div>
  </div>
</template>
<script>
import { ref, onMounted, reactive, computed, onBeforeUnmount } from 'vue'
import { useRouter } from 'vue-router'
import html2canvas from 'html2canvas-pro'
export default {
  name: 'RegisterView',
  setup() {
    const router = useRouter()
    const fullName = ref('')
    const rpgCard = ref(null)
    const submitting = ref(false)
    const error = ref('')
    const avatarDataUrl = ref('')
    const classes = [
      { name: 'Backend Mage' },
      { name: 'Frontend Rogue' },
      { name: 'Cloud Paladin' },
      { name: 'AI Alchemist' }
    ]
    const form = reactive({
      email: '',
      nama_panggilan: '',
      gender: '',
      kota: '',
      provinsi: '',
      pekerjaan: '',
      bio: '',
      rpg_class: 'Cloud Paladin',
      rpg_stats: {
        coding: 75,
        bugs: 75,
        architecture: 75,
        cloud: 75
      },
      linkedin: '',
      background_it: '',
      motivasi: ''
    })
    const provinces = ref([])
    const regencies = ref([])
    const selectedProvince = ref(null)
    const provinceSearch = ref('')
    const regencySearch = ref('')
    const provinceOpen = ref(false)
    const regencyOpen = ref(false)
    const provinceDropdownRef = ref(null)
    const regencyDropdownRef = ref(null)
    const filteredProvinces = computed(() => {
      const q = provinceSearch.value.toLowerCase()
      return q ? provinces.value.filter(p => p.name.toLowerCase().includes(q)) : provinces.value
    })
    const filteredRegencies = computed(() => {
      const q = regencySearch.value.toLowerCase()
      return q ? regencies.value.filter(r => r.name.toLowerCase().includes(q)) : regencies.value
    })
    const toggleProvinceDropdown = () => {
      provinceOpen.value = !provinceOpen.value
      regencyOpen.value = false
      provinceSearch.value = ''
    }
    const toggleRegencyDropdown = () => {
      if (!selectedProvince.value) return
      regencyOpen.value = !regencyOpen.value
      provinceOpen.value = false
      regencySearch.value = ''
    }
    const selectProvince = async (p) => {
      selectedProvince.value = p
      form.provinsi = p.name
      form.kota = ''
      provinceOpen.value = false
      provinceSearch.value = ''
      const fetchRegencies = async (provinceCode) => {
        if (!provinceCode) return
        try {
          const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
          const res = await fetch(`${apiUrl}/api/wilayah/regencies/${provinceCode}`)
          const data = await res.json()
          regencies.value = data.data || []
        } catch {
          regencies.value = []
        }
      }
      fetchRegencies(p.code)
    }
    const selectRegency = (r) => {
      form.kota = r.name
      regencyOpen.value = false
      regencySearch.value = ''
    }
    const handleClickOutside = (e) => {
      if (provinceDropdownRef.value && !provinceDropdownRef.value.contains(e.target)) {
        provinceOpen.value = false
      }
      if (regencyDropdownRef.value && !regencyDropdownRef.value.contains(e.target)) {
        regencyOpen.value = false
      }
    }
    const fetchProvinces = async () => {
      try {
        const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
        const res = await fetch(`${apiUrl}/api/wilayah/provinces`)
        const data = await res.json()
        provinces.value = data.data || []
      } catch {
        provinces.value = []
      }
    }
    onMounted(async () => {
      const email = localStorage.getItem('user_email')
      if (!email) {
        router.push('/')
        return
      }
      form.email = email
      try {
        const response = await fetch(`http://localhost:8080/api/auth?email=${encodeURIComponent(email)}`)
        if (response.ok) {
          const data = await response.json()
          fullName.value = data.participant.nama_lengkap
          if (data.is_registered) {
            router.push(`/v1/share/${data.participant.id}`)
          }
        }
      } catch (err) {
        console.error('Gagal mengambil detail email', err)
      }
      try {
        const pngRes = await fetch(`http://localhost:8080/api/avatar?seed=${encodeURIComponent(email)}`)
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
      fetchProvinces()
      document.addEventListener('click', handleClickOutside)
      generateRandomStats()
    })
    onBeforeUnmount(() => {
      document.removeEventListener('click', handleClickOutside)
    })
    const clamp = (val, min = 55, max = 99) => Math.min(max, Math.max(min, Math.round(val)))
    const noise = () => Math.floor(Math.random() * 11) - 5 
    const calculateAdaptiveStats = () => {
      const isIT = form.background_it === 'IT'
      const cls = form.rpg_class
      const hasLinkedIn = !!form.linkedin
      const bioLen = (form.bio || '').trim().length      
      const motivasiLen = (form.motivasi || '').trim().length 
      let coding = 55
      if (isIT) coding += 20
      if (cls === 'Backend Mage') coding += 15
      else if (cls === 'Frontend Rogue') coding += 12
      else if (cls === 'AI Alchemist') coding += 10
      else if (cls === 'Cloud Paladin') coding += 5
      coding += Math.floor((bioLen / 150) * 8) 
      coding += noise()
      let bugs = 55
      if (isIT) bugs += 15
      if (cls === 'Backend Mage') bugs += 18
      else if (cls === 'Frontend Rogue') bugs += 10
      else if (cls === 'Cloud Paladin') bugs += 8
      else if (cls === 'AI Alchemist') bugs += 12
      bugs += Math.floor((motivasiLen / 150) * 10) 
      bugs += noise()
      let architecture = 55
      if (isIT) architecture += 10
      if (cls === 'Cloud Paladin') architecture += 20
      else if (cls === 'Backend Mage') architecture += 18
      else if (cls === 'AI Alchemist') architecture += 12
      else if (cls === 'Frontend Rogue') architecture += 8
      if (hasLinkedIn) architecture += 5
      architecture += Math.floor(((bioLen + motivasiLen) / 300) * 6) 
      architecture += noise()
      let cloud = 55
      if (isIT) cloud += 10
      if (cls === 'Cloud Paladin') cloud += 25
      else if (cls === 'AI Alchemist') cloud += 20
      else if (cls === 'Backend Mage') cloud += 10
      else if (cls === 'Frontend Rogue') cloud += 6
      if (hasLinkedIn) cloud += 3
      cloud += noise()
      form.rpg_stats.coding       = clamp(coding)
      form.rpg_stats.bugs         = clamp(bugs)
      form.rpg_stats.architecture = clamp(architecture)
      form.rpg_stats.cloud        = clamp(cloud)
    }
    const generateRandomStats = calculateAdaptiveStats
    const cleanLinkedInUrl = (url) => {
      if (!url) return ''
      return url.replace(/https?:\/\/(www\.)?linkedin\.com\/in\//, '')
    }
    const handleSubmit = async () => {
      submitting.value = true
      error.value = ''
      if (form.linkedin) {
        const linkedInRegex = /^https:\/\/(www\.)?linkedin\.com\/in\//
        if (!linkedInRegex.test(form.linkedin)) {
          error.value = 'URL LinkedIn tidak valid. Harap gunakan format https://linkedin.com/in/username'
          submitting.value = false
          return
        }
      }
      try {
        const token = localStorage.getItem('jwt_token') || ''
        const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
        const response = await fetch(`${apiUrl}/api/profile`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${token}`
          },
          body: JSON.stringify(form)
        })
        const data = await response.json()
        if (!response.ok) {
          error.value = data.error || 'Gagal menyimpan profil.'
          submitting.value = false
          return
        }
        const participantId = data.participant.id
        await new Promise(resolve => setTimeout(resolve, 300))
        if (!rpgCard.value) {
          error.value = 'Elemen kartu pratinjau tidak ditemukan.'
          submitting.value = false
          return
        }
        const canvas = await html2canvas(rpgCard.value, {
          scale: 3, 
          useCORS: true,
          backgroundColor: null 
        })
        canvas.toBlob(async (blob) => {
          if (!blob) {
            error.value = 'Gagal mengonversi kartu ke gambar.'
            submitting.value = false
            return
          }
          const formData = new FormData()
          formData.append('image', blob, `card_${participantId}.png`)
          const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
          const uploadResponse = await fetch(`${apiUrl}/api/profile/${participantId}/upload`, {
            method: 'POST',
            headers: {
              'Authorization': `Bearer ${token}`
            },
            body: formData
          })
          if (uploadResponse.ok) {
            const slugify = (text) => {
              if (!text) return ''
              return text.toString().toLowerCase()
                .replace(/\s+/g, '-')
                .replace(/[^\w\-]+/g, '')
                .replace(/\-\-+/g, '-')
                .replace(/^-+/, '')
                .replace(/-+$/, '')
            }
            const nameSlug = slugify(data.participant.nama_lengkap || '')
            router.push(`/v1/share/${participantId}/${nameSlug}`)
          } else {
            const uploadData = await uploadResponse.json()
            error.value = uploadData.error || 'Gagal mengunggah kartu ke server.'
            submitting.value = false
          }
        }, 'image/png')
      } catch (err) {
        error.value = 'Terjadi kesalahan saat menyimpan profil: ' + err.message
        console.error(err)
        submitting.value = false
      }
    }
    return {
      fullName,
      classes,
      form,
      rpgCard,
      submitting,
      error,
      generateRandomStats,
      calculateAdaptiveStats,
      handleSubmit,
      avatarDataUrl,
      cleanLinkedInUrl,
      provinceSearch,
      regencySearch,
      provinceOpen,
      regencyOpen,
      provinceDropdownRef,
      regencyDropdownRef,
      selectedProvince,
      filteredProvinces,
      filteredRegencies,
      toggleProvinceDropdown,
      toggleRegencyDropdown,
      selectProvince,
      selectRegency
    }
  }
}
</script>

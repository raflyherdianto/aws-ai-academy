<template>
  <div class="min-h-screen flex items-center justify-center px-4 py-12">
    <div class="w-full max-w-lg mx-auto">
      <transition name="fade" mode="out-in">

        <!-- ===== STEP 1 ===== -->
        <div v-if="step === 1" key="step1" class="glass-container p-8 relative overflow-hidden text-center">
          <div class="absolute -top-16 -left-16 w-40 h-40 bg-green-500/10 rounded-full blur-3xl pointer-events-none"></div>
          <div class="absolute -bottom-16 -right-16 w-40 h-40 bg-aws-orange/10 rounded-full blur-3xl pointer-events-none"></div>

          <div class="inline-flex items-center gap-2 px-3 py-1 bg-green-500/10 border border-green-500/30 rounded-full mb-5">
            <span class="text-green-400 text-xs font-bold uppercase tracking-wider">🤝 Komitmen Peserta</span>
          </div>
          <h1 class="text-2xl sm:text-3xl font-extrabold text-white mb-2 leading-tight">
            Bantu Fasilitator<br /><span class="text-green-400">Mencapai Target!</span>
          </h1>
          <p class="text-slate-400 text-sm mb-6">
            Berjanji dan berkomitmen untuk menyelesaikan seluruh kelas tepat waktu dan <strong class="text-rose-400">tidak ghosting</strong>
          </p>

          <div class="relative mx-auto mb-6 w-full max-w-xs rounded-2xl overflow-hidden border border-white/10 shadow-2xl shadow-black/50">
            <img :src="img1" alt="Komitmen" class="w-full object-cover transition-opacity duration-300" :class="img1Loaded ? 'opacity-100' : 'opacity-0'" @load="img1Loaded = true" />
            <div v-if="!img1Loaded" class="absolute inset-0 bg-slate-800 animate-pulse flex items-center justify-center"><span class="text-4xl">🤝</span></div>
          </div>

          <div class="mb-7 text-left space-y-2 px-2">
            <div class="flex items-start gap-2 text-sm text-slate-300"><span class="text-green-400 mt-0.5">✅</span><span>Selesaikan seluruh kelas tepat waktu</span></div>
            <div class="flex items-start gap-2 text-sm text-slate-300"><span class="text-green-400 mt-0.5">✅</span><span>Aktif di setiap sesi dan tidak ghosting</span></div>
            <div class="flex items-start gap-2 text-sm text-slate-300"><span class="text-green-400 mt-0.5">✅</span><span>Bantu fasilitator mencapai target kelulusan</span></div>
          </div>

          <p v-if="error" class="mb-4 text-rose-400 text-xs">{{ error }}</p>

          <div class="grid grid-cols-2 gap-4">
            <!-- YES1 → Step 3 (hadiah spesial) -->
            <button id="btn-yes1" @click="step = 3" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-green-500/50 hover:border-green-400 transition-all duration-200 cursor-pointer shadow-lg shadow-green-500/20">
              <img :src="yes1" alt="Ya" class="w-full object-cover" />
            </button>
            <!-- NO1 → Step 2 (alam semesta) -->
            <button id="btn-no1" @click="step = 2" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-slate-600/50 hover:border-slate-400 transition-all duration-200 cursor-pointer shadow-lg">
              <img :src="no1" alt="Tidak" class="w-full object-cover" />
            </button>
          </div>
          <p class="mt-4 text-slate-600 text-xs">Pilihan Anda tidak mempengaruhi kartu developer Anda.</p>
        </div>

        <!-- ===== STEP 2 — Alam Semesta ===== -->
        <div v-else-if="step === 2" key="step2" class="glass-container p-8 relative overflow-hidden text-center">
          <div class="absolute -top-16 -right-16 w-40 h-40 bg-amber-500/10 rounded-full blur-3xl pointer-events-none"></div>
          <div class="absolute -bottom-16 -left-16 w-40 h-40 bg-purple-500/10 rounded-full blur-3xl pointer-events-none"></div>

          <div class="inline-flex items-center gap-2 px-3 py-1 bg-amber-500/10 border border-amber-500/30 rounded-full mb-5">
            <span class="text-amber-400 text-xs font-bold uppercase tracking-wider">🌌 Renungkan</span>
          </div>
          <h1 class="text-2xl sm:text-3xl font-extrabold text-white mb-2 leading-tight">
            Pikir Ulang<br /><span class="text-amber-400">Keputusanmu!</span>
          </h1>
          <p class="text-slate-300 text-sm mb-6 leading-relaxed">
            Dengan <strong class="text-amber-300">rela membantu fasilitator</strong>, alam semesta akan senantiasa membantu kehidupan Anda. ✨
          </p>

          <div class="relative mx-auto mb-6 w-full max-w-xs rounded-2xl overflow-hidden border border-white/10 shadow-2xl shadow-black/50">
            <img :src="img2" alt="Alam semesta" class="w-full object-cover transition-opacity duration-300" :class="img2Loaded ? 'opacity-100' : 'opacity-0'" @load="img2Loaded = true" />
            <div v-if="!img2Loaded" class="absolute inset-0 bg-slate-800 animate-pulse flex items-center justify-center"><span class="text-4xl">🌌</span></div>
          </div>

          <p v-if="error" class="mb-4 text-rose-400 text-xs">{{ error }}</p>

          <div class="grid grid-cols-2 gap-4">
            <!-- YES2 → step = 3 -->
            <button id="btn-yes2" @click="step = 3" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-amber-500/50 hover:border-amber-400 transition-all duration-200 cursor-pointer shadow-lg shadow-amber-500/20">
              <img :src="yes2" alt="Ya" class="w-full object-cover" />
            </button>
            <!-- NO2 → step = 4 -->
            <button id="btn-no2" @click="step = 4" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-slate-600/50 hover:border-slate-400 transition-all duration-200 cursor-pointer shadow-lg">
              <img :src="no2" alt="Tidak" class="w-full object-cover" />
            </button>
          </div>
          <button @click="step = 1" class="mt-4 text-slate-600 text-xs hover:text-slate-400 transition-colors cursor-pointer">← Kembali</button>
        </div>

        <!-- ===== STEP 3 — Hadiah Spesial ===== -->
        <div v-else-if="step === 3" key="step3" class="glass-container p-8 relative overflow-hidden text-center">
          <div class="absolute -top-16 -left-16 w-40 h-40 bg-yellow-500/10 rounded-full blur-3xl pointer-events-none"></div>
          <div class="absolute -bottom-16 -right-16 w-40 h-40 bg-pink-500/10 rounded-full blur-3xl pointer-events-none"></div>

          <div class="inline-flex items-center gap-2 px-3 py-1 bg-yellow-500/10 border border-yellow-500/30 rounded-full mb-5">
            <span class="text-yellow-400 text-xs font-bold uppercase tracking-wider">🎁 Hadiah Spesial</span>
          </div>
          <h1 class="text-2xl sm:text-3xl font-extrabold text-white mb-2 leading-tight">
            Ada Sesuatu<br /><span class="text-yellow-400">Untukmu!</span>
          </h1>
          <p class="text-slate-300 text-sm mb-6 leading-relaxed">
            Dengan <strong class="text-yellow-300">membantu fasilitator</strong>, Anda akan berkesempatan mendapatkan <strong class="text-yellow-300">hadiah spesial</strong>! 🎁
          </p>

          <div class="relative mx-auto mb-6 w-full max-w-xs rounded-2xl overflow-hidden border border-white/10 shadow-2xl shadow-black/50">
            <img :src="img3" alt="Hadiah Spesial" class="w-full object-cover transition-opacity duration-300" :class="img3Loaded ? 'opacity-100' : 'opacity-0'" @load="img3Loaded = true" />
            <div v-if="!img3Loaded" class="absolute inset-0 bg-slate-800 animate-pulse flex items-center justify-center"><span class="text-4xl">🎁</span></div>
          </div>

          <p v-if="error" class="mb-4 text-rose-400 text-xs">{{ error }}</p>

          <div class="grid grid-cols-2 gap-4">
            <!-- YES3 → step = 5 -->
            <button id="btn-yes3" @click="step = 5" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-yellow-500/50 hover:border-yellow-400 transition-all duration-200 cursor-pointer shadow-lg shadow-yellow-500/20">
              <img :src="yes3" alt="Ya" class="w-full object-cover" />
            </button>
            <!-- NO3 → step = 5 -->
            <button id="btn-no3" @click="step = 5" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-slate-600/50 hover:border-slate-400 transition-all duration-200 cursor-pointer shadow-lg">
              <img :src="no3" alt="Tidak" class="w-full object-cover" />
            </button>
          </div>
          <button @click="step = 1" class="mt-4 text-slate-600 text-xs hover:text-slate-400 transition-colors cursor-pointer">← Kembali</button>
        </div>

        <!-- ===== STEP 4 — Permohonan ===== -->
        <div v-else-if="step === 4" key="step4" class="glass-container p-8 relative overflow-hidden text-center">
          <div class="absolute -top-16 -right-16 w-40 h-40 bg-indigo-500/10 rounded-full blur-3xl pointer-events-none"></div>
          <div class="absolute -bottom-16 -left-16 w-40 h-40 bg-teal-500/10 rounded-full blur-3xl pointer-events-none"></div>

          <div class="inline-flex items-center gap-2 px-3 py-1 bg-indigo-500/10 border border-indigo-500/30 rounded-full mb-5">
            <span class="text-indigo-400 text-xs font-bold uppercase tracking-wider">🙏 Permohonan</span>
          </div>
          <h1 class="text-2xl sm:text-3xl font-extrabold text-white mb-2 leading-tight">
            Tolonglah<br /><span class="text-indigo-400">Saya!</span>
          </h1>
          <p class="text-slate-300 text-sm mb-6 leading-relaxed">
            Saya mohon Anda untuk <strong class="text-indigo-300">berkomitmen membantu fasilitator</strong> demi kelancaran program.
          </p>

          <div class="relative mx-auto mb-6 w-full max-w-xs rounded-2xl overflow-hidden border border-white/10 shadow-2xl shadow-black/50">
            <img :src="img4" alt="Mohon Komitmen" class="w-full object-cover transition-opacity duration-300" :class="img4Loaded ? 'opacity-100' : 'opacity-0'" @load="img4Loaded = true" />
            <div v-if="!img4Loaded" class="absolute inset-0 bg-slate-800 animate-pulse flex items-center justify-center"><span class="text-4xl">🙏</span></div>
          </div>

          <p v-if="error" class="mb-4 text-rose-400 text-xs">{{ error }}</p>

          <div class="grid grid-cols-2 gap-4">
            <!-- YES4 → redirectOnly (jangan atur commitment = 1) -->
            <button id="btn-yes4" @click="step = 3" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-indigo-500/50 hover:border-indigo-400 transition-all duration-200 cursor-pointer shadow-lg shadow-indigo-500/20">
              <img :src="yes4" alt="Ya" class="w-full object-cover" />
            </button>
            <!-- NO4 → step = 3 -->
            <button id="btn-no4" @click="step = 3" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-slate-600/50 hover:border-slate-400 transition-all duration-200 cursor-pointer shadow-lg">
              <img :src="no4" alt="Tidak" class="w-full object-cover" />
            </button>
          </div>
          <button @click="step = 2" class="mt-4 text-slate-600 text-xs hover:text-slate-400 transition-colors cursor-pointer">← Kembali</button>
        </div>

        <!-- ===== STEP 5 — Senyuman Imut ===== -->
        <div v-else-if="step === 5" key="step5" class="glass-container p-8 relative overflow-hidden text-center">
          <div class="absolute -top-16 -right-16 w-40 h-40 bg-pink-500/10 rounded-full blur-3xl pointer-events-none"></div>
          <div class="absolute -bottom-16 -left-16 w-40 h-40 bg-purple-500/10 rounded-full blur-3xl pointer-events-none"></div>

          <div class="inline-flex items-center gap-2 px-3 py-1 bg-pink-500/10 border border-pink-500/30 rounded-full mb-5">
            <span class="text-pink-400 text-xs font-bold uppercase tracking-wider">😊 Senyuman</span>
          </div>
          <h1 class="text-2xl sm:text-3xl font-extrabold text-white mb-2 leading-tight">
            Hadiah<br /><span class="text-pink-400">Senyuman Imut</span>
          </h1>
          <p class="text-slate-300 text-sm mb-6 leading-relaxed">
            <strong class="text-pink-300">Hadiah senyuman imut untukmu!</strong> 😊
          </p>

          <div class="relative mx-auto mb-6 w-full max-w-xs rounded-2xl overflow-hidden border border-white/10 shadow-2xl shadow-black/50">
            <img :src="img5" alt="Senyuman Imut" class="w-full object-cover transition-opacity duration-300" :class="img5Loaded ? 'opacity-100' : 'opacity-0'" @load="img5Loaded = true" />
            <div v-if="!img5Loaded" class="absolute inset-0 bg-slate-800 animate-pulse flex items-center justify-center"><span class="text-4xl">😊</span></div>
          </div>

          <p v-if="error" class="mb-4 text-rose-400 text-xs">{{ error }}</p>

          <div class="grid grid-cols-2 gap-4">
            <!-- YES5 → step = 6 -->
            <button id="btn-yes5" @click="step = 6" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-pink-500/50 hover:border-pink-400 transition-all duration-200 cursor-pointer shadow-lg shadow-pink-500/20">
              <img :src="yes5" alt="Ya" class="w-full object-cover" />
            </button>
            <!-- NO5 → step = 6 -->
            <button id="btn-no5" @click="step = 6" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-slate-600/50 hover:border-slate-400 transition-all duration-200 cursor-pointer shadow-lg">
              <img :src="no5" alt="Tidak" class="w-full object-cover" />
            </button>
          </div>
          <button @click="step = 3" class="mt-4 text-slate-600 text-xs hover:text-slate-400 transition-colors cursor-pointer">← Kembali</button>
        </div>

        <!-- ===== STEP 6 — Grand Prize ===== -->
        <div v-else-if="step === 6" key="step6" class="glass-container p-8 relative overflow-hidden text-center">
          <div class="absolute -top-16 -right-16 w-40 h-40 bg-yellow-500/10 rounded-full blur-3xl pointer-events-none"></div>
          <div class="absolute -bottom-16 -left-16 w-40 h-40 bg-orange-500/10 rounded-full blur-3xl pointer-events-none"></div>

          <div class="inline-flex items-center gap-2 px-3 py-1 bg-yellow-500/10 border border-yellow-500/30 rounded-full mb-5">
            <span class="text-yellow-400 text-xs font-bold uppercase tracking-wider">🏆 Grand Prize</span>
          </div>
          <h1 class="text-2xl sm:text-3xl font-extrabold text-white mb-2 leading-tight">
            Hadiah<br /><span class="text-yellow-400">Utama</span>
          </h1>
          <p class="text-slate-300 text-sm mb-6 leading-relaxed">
            <strong class="text-yellow-300">Hadiah grand prize</strong> jika Anda benar-benar <strong class="text-yellow-300">berkomitmen membantu fasilitator</strong>! 🏆
          </p>

          <div class="relative mx-auto mb-6 w-full max-w-xs rounded-2xl overflow-hidden border border-white/10 shadow-2xl shadow-black/50">
            <img :src="img6" alt="Grand Prize" class="w-full object-cover transition-opacity duration-300" :class="img6Loaded ? 'opacity-100' : 'opacity-0'" @load="img6Loaded = true" />
            <div v-if="!img6Loaded" class="absolute inset-0 bg-slate-800 animate-pulse flex items-center justify-center"><span class="text-4xl">🏆</span></div>
          </div>

          <p v-if="error" class="mb-4 text-rose-400 text-xs">{{ error }}</p>

          <div class="grid grid-cols-2 gap-4">
            <!-- YES6 → step = 7 -->
            <button id="btn-yes6" @click="step = 7" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-yellow-500/50 hover:border-yellow-400 transition-all duration-200 cursor-pointer shadow-lg shadow-yellow-500/20">
              <img :src="yes6" alt="Ya" class="w-full object-cover" />
            </button>
            <!-- NO6 → step = 7 -->
            <button id="btn-no6" @click="step = 7" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-slate-600/50 hover:border-slate-400 transition-all duration-200 cursor-pointer shadow-lg">
              <img :src="no6" alt="Tidak" class="w-full object-cover" />
            </button>
          </div>
          <button @click="step = 5" class="mt-4 text-slate-600 text-xs hover:text-slate-400 transition-colors cursor-pointer">← Kembali</button>
        </div>

        <!-- ===== STEP 7 — Permohonan Terakhir ===== -->
        <div v-else-if="step === 7" key="step7" class="glass-container p-8 relative overflow-hidden text-center">
          <div class="absolute -top-16 -right-16 w-40 h-40 bg-red-500/10 rounded-full blur-3xl pointer-events-none"></div>
          <div class="absolute -bottom-16 -left-16 w-40 h-40 bg-indigo-500/10 rounded-full blur-3xl pointer-events-none"></div>

          <div class="inline-flex items-center gap-2 px-3 py-1 bg-red-500/10 border border-red-500/30 rounded-full mb-5">
            <span class="text-red-400 text-xs font-bold uppercase tracking-wider">🙏 Tolonglah</span>
          </div>
          <h1 class="text-2xl sm:text-3xl font-extrabold text-white mb-2 leading-tight">
            Sekali Lagi,<br /><span class="text-red-400">Kumohon!</span>
          </h1>
          <p class="text-slate-300 text-sm mb-6 leading-relaxed">
            Apakah Anda <strong class="text-red-300">yakin ingin</strong> membantu fasilitator?
          </p>

          <div class="relative mx-auto mb-6 w-full max-w-xs rounded-2xl overflow-hidden border border-white/10 shadow-2xl shadow-black/50">
            <img :src="img7" alt="Permohonan Terakhir" class="w-full object-cover transition-opacity duration-300" :class="img7Loaded ? 'opacity-100' : 'opacity-0'" @load="img7Loaded = true" />
            <div v-if="!img7Loaded" class="absolute inset-0 bg-slate-800 animate-pulse flex items-center justify-center"><span class="text-4xl">🥺</span></div>
          </div>

          <p v-if="error" class="mb-4 text-rose-400 text-xs">{{ error }}</p>

          <div class="grid grid-cols-2 gap-4">
            <!-- YES7 → step = 8 -->
            <button id="btn-yes7" @click="step = 8" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-red-500/50 hover:border-red-400 transition-all duration-200 cursor-pointer shadow-lg shadow-red-500/20">
              <img :src="yes7" alt="Ya" class="w-full object-cover" />
            </button>
            <!-- NO7 → step = 8 -->
            <button id="btn-no7" @click="step = 8" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-slate-600/50 hover:border-slate-400 transition-all duration-200 cursor-pointer shadow-lg">
              <img :src="no7" alt="Tidak" class="w-full object-cover" />
            </button>
          </div>
          <button @click="step = 6" class="mt-4 text-slate-600 text-xs hover:text-slate-400 transition-colors cursor-pointer">← Kembali</button>
        </div>

        <!-- ===== STEP 8 — Komitmen Akhir ===== -->
        <div v-else-if="step === 8" key="step8" class="glass-container p-8 relative overflow-hidden text-center">
          <div class="absolute -top-16 -right-16 w-40 h-40 bg-emerald-500/10 rounded-full blur-3xl pointer-events-none"></div>
          <div class="absolute -bottom-16 -left-16 w-40 h-40 bg-cyan-500/10 rounded-full blur-3xl pointer-events-none"></div>

          <div class="inline-flex items-center gap-2 px-3 py-1 bg-emerald-500/10 border border-emerald-500/30 rounded-full mb-5">
            <span class="text-emerald-400 text-xs font-bold uppercase tracking-wider">🤝 Perjanjian</span>
          </div>
          <h1 class="text-2xl sm:text-3xl font-extrabold text-white mb-2 leading-tight">
            Janji<br /><span class="text-emerald-400">Komitmen</span>
          </h1>
          <p class="text-slate-300 text-sm mb-6 leading-relaxed">
            <strong class="text-emerald-300">Dengan ini saya menyatakan berjanji dan berkomitmen</strong> untuk membantu fasilitator dengan menyelesaikan seluruh kelas tepat waktu dan tidak ghosting selama program berlangsung. 🤝
          </p>

          <div class="relative mx-auto mb-6 w-full max-w-xs rounded-2xl overflow-hidden border border-white/10 shadow-2xl shadow-black/50">
            <img :src="img8" alt="Janji Komitmen" class="w-full object-cover transition-opacity duration-300" :class="img8Loaded ? 'opacity-100' : 'opacity-0'" @load="img8Loaded = true" />
            <div v-if="!img8Loaded" class="absolute inset-0 bg-slate-800 animate-pulse flex items-center justify-center"><span class="text-4xl">🤝</span></div>
          </div>

          <p v-if="error" class="mb-4 text-rose-400 text-xs">{{ error }}</p>

          <div class="grid grid-cols-2 gap-4">
            <!-- YES8 → commitment=1 -->
            <button id="btn-yes8" @click="handleCommitment(1, 'yes8')" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-emerald-500/50 hover:border-emerald-400 transition-all duration-200 cursor-pointer shadow-lg shadow-emerald-500/20">
              <img :src="yes8" alt="Ya" class="w-full object-cover" />
              <span v-if="loading && selectedChoice === 'yes8'" class="absolute inset-0 flex items-center justify-center bg-emerald-900/70 rounded-xl">
                <span class="w-6 h-6 border-2 border-white border-t-transparent rounded-full animate-spin"></span>
              </span>
            </button>
            <!-- NO8 → commitment=1 -->
            <button id="btn-no8" @click="handleCommitment(1, 'no8')" :disabled="loading"
              class="group relative rounded-xl overflow-hidden border-2 border-emerald-500/50 hover:border-emerald-400 transition-all duration-200 cursor-pointer shadow-lg shadow-emerald-500/20">
              <img :src="no8" alt="Tidak" class="w-full object-cover" />
              <span v-if="loading && selectedChoice === 'no8'" class="absolute inset-0 flex items-center justify-center bg-emerald-900/70 rounded-xl">
                <span class="w-6 h-6 border-2 border-white border-t-transparent rounded-full animate-spin"></span>
              </span>
            </button>
          </div>
          <button @click="step = 7" class="mt-4 text-slate-600 text-xs hover:text-slate-400 transition-colors cursor-pointer">← Kembali</button>
        </div>
      </transition>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import img1 from '../assets/img1.jpg'
import img2 from '../assets/img2.webp'
import img3 from '../assets/img3.jpg'
import img4 from '../assets/img4.jpg'
import img5 from '../assets/img5.webp'
import img6 from '../assets/img6.jpg'
import img7 from '../assets/img7.webp'
import img8 from '../assets/img8.webp'
import yes1 from '../assets/yes1.webp'
import yes2 from '../assets/yes2.jpg'
import yes3 from '../assets/yes3.webp'
import yes4 from '../assets/yes4.webp'
import yes5 from '../assets/yes5.jpg'
import yes6 from '../assets/yes6.jpg'
import yes7 from '../assets/yes7.webp'
import yes8 from '../assets/yes8.jpg'
import no1 from '../assets/no1.jpg'
import no2 from '../assets/no2.webp'
import no3 from '../assets/no3.jpg'
import no4 from '../assets/no4.webp'
import no5 from '../assets/no5.webp'
import no6 from '../assets/no6.webp'
import no7 from '../assets/no7.webp'
import no8 from '../assets/no8.webp'

export default {
  name: 'CommitmentView',
  setup() {
    const router = useRouter()
    const step = ref(1)
    const loading = ref(false)
    const error = ref('')
    const selectedChoice = ref(null)
    const img1Loaded = ref(false)
    const img2Loaded = ref(false)
    const img3Loaded = ref(false)
    const img4Loaded = ref(false)
    const img5Loaded = ref(false)
    const img6Loaded = ref(false)
    const img7Loaded = ref(false)
    const img8Loaded = ref(false)
    const participantId = ref(null)
    const nameSlug = ref('')

    onMounted(() => {
      const id = localStorage.getItem('pending_participant_id')
      const slug = localStorage.getItem('pending_name_slug')
      if (!id) { router.push('/'); return }
      participantId.value = id
      nameSlug.value = slug || ''
    })

    const doRedirect = () => {
      localStorage.removeItem('pending_participant_id')
      localStorage.removeItem('pending_name_slug')
      router.push(`/share-card/${participantId.value}/${nameSlug.value}`)
    }

    // Redirect tanpa set commitment (yes1, yes3)
    const redirectOnly = (choice) => {
      if (loading.value) return
      selectedChoice.value = choice
      loading.value = true
      doRedirect()
    }

    // Set commitment ke API lalu redirect (yes2→1, no2→-1, no3→-1)
    const handleCommitment = async (value, choice) => {
      if (loading.value) return
      selectedChoice.value = choice
      loading.value = true
      error.value = ''
      try {
        const token = localStorage.getItem('jwt_token') || ''
        const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
        await fetch(`${apiUrl}/api/profile/${participantId.value}/commitment`, {
          method: 'PATCH',
          headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
          body: JSON.stringify({ commitment: value })
        })
      } catch (err) {
        console.warn('Commitment save failed (non-blocking):', err)
      }
      doRedirect()
    }

    return {
      step, loading, error, selectedChoice,
      img1Loaded, img2Loaded, img3Loaded, img4Loaded, img5Loaded, img6Loaded, img7Loaded, img8Loaded,
      img1, img2, img3, img4, img5, img6, img7, img8,
      yes1, yes2, yes3, yes4, yes5, yes6, yes7, yes8,
      no1, no2, no3, no4, no5, no6, no7, no8,
      redirectOnly, handleCommitment
    }
  }
}
</script>

<style scoped>
.fade-enter-active, .fade-leave-active { transition: opacity 0.25s ease, transform 0.25s ease; }
.fade-enter-from { opacity: 0; transform: translateX(20px); }
.fade-leave-to  { opacity: 0; transform: translateX(-20px); }
</style>

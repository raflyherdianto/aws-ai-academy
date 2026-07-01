<template>
  <div class="w-full max-w-xl mx-auto flex flex-col gap-6">
    <div v-if="gameState === 'start'" class="glass-container p-8 text-center relative overflow-hidden">
      <div class="absolute -top-10 -right-10 w-32 h-32 bg-aws-orange/10 rounded-full blur-xl animate-pulse"></div>
      <h2 class="text-3xl font-extrabold text-white mb-2">Guess Who? 🔍</h2>
      <p class="text-slate-400 text-sm mb-6">
        Game Tebak Profil Peserta AWS AI Academy! Uji seberapa baik Anda mengenal rekan-rekan satu cohort melalui stats, bio, dan asal domisili mereka.
      </p>
      <button 
        @click="startGame"
        class="px-8 py-3.5 bg-aws-orange hover:bg-aws-orange/90 text-slate-950 font-bold rounded-xl text-sm transition-all shadow-lg shadow-aws-orange/20 cursor-pointer"
      >
        Mulai Game
      </button>
    </div>
    <div v-else-if="gameState === 'gameover'" class="glass-container p-8 text-center relative overflow-hidden">
      <h2 class="text-2xl font-bold text-white mb-2">Game Selesai! 🏆</h2>
      <p class="text-slate-400 text-sm mb-6">Terima kasih sudah bermain ice-breaking ini.</p>
      <div class="p-6 bg-white/5 border border-white/10 rounded-2xl mb-8 max-w-xs mx-auto">
        <span class="text-xs text-slate-500 block uppercase tracking-wider font-semibold">Skor Akhir Anda</span>
        <span class="text-4xl font-extrabold text-aws-orange font-mono block mt-1">{{ totalScore }} Poin</span>
        <span class="text-[10px] text-slate-400 block mt-2">Menjawab {{ correctAnswers }} dari 5 pertanyaan dengan benar</span>
      </div>
      <div class="flex gap-4 justify-center">
        <button 
          @click="startGame"
          class="px-6 py-3 bg-dicoding-blue hover:bg-dicoding-blue/90 text-white font-bold rounded-xl text-xs cursor-pointer transition-all"
        >
          Main Lagi
        </button>
        <router-link 
          to="/"
          class="px-6 py-3 bg-white/5 hover:bg-white/10 border border-white/10 text-white font-bold rounded-xl text-xs transition-all"
        >
          Ke Beranda
        </router-link>
      </div>
    </div>
    <div v-else-if="gameState === 'playing'" class="space-y-6">
      <div class="flex justify-between items-center px-2">
        <span class="text-xs font-semibold text-slate-400">Pertanyaan {{ round }} dari 5</span>
        <span class="text-xs font-mono font-bold text-aws-orange bg-aws-orange/10 px-2.5 py-1 rounded-full border border-aws-orange/20">
          Skor: {{ totalScore }}
        </span>
      </div>
      <div class="glass-container p-6 relative overflow-hidden">
        <div class="absolute -top-10 -right-10 w-24 h-24 bg-dicoding-blue/10 rounded-full blur-xl"></div>
        <h3 class="text-sm font-bold text-slate-300 mb-4 flex items-center gap-1.5">
          <span>🔎</span> Tebak Siapa Saya?
        </h3>
        <div class="space-y-3 min-h-[160px] flex flex-col justify-center">
          <transition-group name="list">
            <div 
              v-for="(clue, index) in visibleClues" 
              :key="index"
              class="p-3 bg-white/5 border border-white/5 rounded-xl text-xs text-slate-200"
            >
              • {{ clue }}
            </div>
          </transition-group>
        </div>
        <div class="mt-6 flex justify-between items-center border-t border-white/5 pt-4">
          <span class="text-[10px] text-slate-500 font-mono">
            Sisa Petunjuk: {{ question.clues ? question.clues.length - clueIndex : 0 }}
          </span>
          <button 
            @click="revealNextClue"
            :disabled="clueIndex >= (question.clues ? question.clues.length : 0) || answered"
            class="px-3 py-1.5 bg-white/5 hover:bg-white/10 border border-white/10 text-[10px] font-bold text-white rounded-lg transition-all cursor-pointer disabled:opacity-40 disabled:cursor-not-allowed"
          >
            💡 Minta Petunjuk (+1)
          </button>
        </div>
      </div>
      <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
        <button 
          v-for="choice in question.choices" 
          :key="choice"
          @click="selectAnswer(choice)"
          :disabled="answered"
          class="p-4 rounded-xl border text-left text-sm font-semibold transition-all cursor-pointer flex justify-between items-center"
          :class="getChoiceClass(choice)"
        >
          <span>{{ choice }}</span>
          <span v-if="answered && choice === question.correct" class="text-emerald-400">✓</span>
          <span v-else-if="answered && selectedAnswer === choice && choice !== question.correct" class="text-rose-400">✗</span>
        </button>
      </div>
      <div v-if="answered" class="text-center pt-2">
        <button 
          @click="nextQuestion"
          class="px-8 py-3 bg-aws-orange hover:bg-aws-orange/90 text-slate-950 font-bold rounded-xl text-sm transition-all cursor-pointer shadow-lg shadow-aws-orange/10"
        >
          {{ round === 5 ? 'Lihat Skor Akhir' : 'Lanjut ke Pertanyaan Berikut' }}
        </button>
      </div>
    </div>
    <div v-else-if="gameState === 'error'" class="glass-container p-8 text-center">
      <p class="text-rose-400 font-semibold mb-4">{{ error }}</p>
      <router-link to="/" class="px-4 py-2 bg-aws-orange hover:bg-aws-orange/90 text-slate-950 font-bold rounded-xl text-sm transition-all">
        Kembali & Buat Kartu Pertama
      </router-link>
    </div>
  </div>
</template>
<script>
import { ref } from 'vue'
export default {
  name: 'GameView',
  setup() {
    const gameState = ref('start') 
    const round = ref(1)
    const totalScore = ref(0)
    const correctAnswers = ref(0)
    const loading = ref(false)
    const error = ref('')
    const question = ref({})
    const visibleClues = ref([])
    const clueIndex = ref(0)
    const answered = ref(false)
    const selectedAnswer = ref('')
    const startGame = () => {
      round.value = 1
      totalScore.value = 0
      correctAnswers.value = 0
      fetchNextQuestion()
    }
    const fetchNextQuestion = async () => {
      loading.value = true
      error.value = ''
      answered.value = false
      selectedAnswer.value = ''
      visibleClues.value = []
      clueIndex.value = 0
      try {
        const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
        const response = await fetch(`${apiUrl}/api/game/quiz`)
        const data = await response.json()
        if (!response.ok) {
          error.value = data.error || 'Gagal memuat kuis.'
          gameState.value = 'error'
          return
        }
        if (data.clues && data.clues.length > 4) {
          data.clues = data.clues.slice(0, 4)
        }
        question.value = data
        gameState.value = 'playing'
        if (data.clues && data.clues.length > 0) {
          revealNextClue()
        }
      } catch (err) {
        error.value = 'Gagal memuat pertanyaan. Pastikan server backend berjalan.'
        gameState.value = 'error'
      } finally {
        loading.value = false
      }
    }
    const revealNextClue = () => {
      if (!question.value.clues) return
      if (clueIndex.value < question.value.clues.length) {
        visibleClues.value.push(question.value.clues[clueIndex.value])
        clueIndex.value++
      }
    }
    const selectAnswer = (choice) => {
      if (answered.value) return
      answered.value = true
      selectedAnswer.value = choice
      if (choice === question.value.correct) {
        correctAnswers.value++
        const pointsAwarded = Math.max(10, 40 - (clueIndex.value * 5))
        totalScore.value += pointsAwarded
      }
    }
    const getChoiceClass = (choice) => {
      if (!answered.value) {
        return 'border-white/10 bg-white/5 text-slate-300 hover:bg-white/10'
      }
      if (choice === question.value.correct) {
        return 'border-emerald-500 bg-emerald-500/10 text-emerald-400'
      }
      if (selectedAnswer.value === choice) {
        return 'border-rose-500 bg-rose-500/10 text-rose-400'
      }
      return 'border-white/5 bg-white/2 text-slate-500 opacity-60'
    }
    const nextQuestion = () => {
      if (round.value >= 5) {
        gameState.value = 'gameover'
      } else {
        round.value++
        fetchNextQuestion()
      }
    }
    return {
      gameState,
      round,
      totalScore,
      correctAnswers,
      loading,
      error,
      question,
      visibleClues,
      clueIndex,
      answered,
      selectedAnswer,
      startGame,
      revealNextClue,
      selectAnswer,
      getChoiceClass,
      nextQuestion
    }
  }
}
</script>
<style>
.list-enter-active,
.list-leave-active {
  transition: all 0.3s ease;
}
.list-enter-from {
  opacity: 0;
  transform: translateY(10px);
}
.list-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>

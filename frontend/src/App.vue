<template>
  <div class="min-h-screen flex flex-col">
    <header class="sticky top-0 z-50 px-4 md:px-6 py-4 bg-slate-950/60 backdrop-blur-md border-b border-white/5">
      <div class="max-w-6xl mx-auto flex flex-col md:flex-row items-center justify-between gap-3 md:gap-0">
        <router-link to="/" class="flex items-center gap-2 md:gap-3">
          <span class="w-2 md:w-2.5 h-2 md:h-2.5 rounded-full bg-aws-orange animate-pulse"></span>
          <span class="font-bold text-base md:text-xl tracking-tight text-white font-mono text-center">
            AWS AI <span class="text-aws-orange">Academy</span> <span class="hidden sm:inline">Networking</span>
          </span>
        </router-link>
        <nav class="flex items-center gap-4 md:gap-6">
          <router-link 
            to="/" 
            class="text-xs md:text-sm font-medium transition-colors hover:text-aws-orange"
            :class="$route.path === '/' || $route.path === '/register' ? 'text-aws-orange' : 'text-slate-300'"
          >
            Buat Kartu
          </router-link>
          <router-link 
            to="/game" 
            class="text-xs md:text-sm font-medium transition-colors hover:text-aws-orange"
            :class="$route.path === '/game' ? 'text-aws-orange' : 'text-slate-300'"
          >
            Tebak Profil (Game)
          </router-link>
        </nav>
      </div>
    </header>
    <main class="flex-grow flex items-center justify-center p-4 md:p-8 max-w-6xl mx-auto w-full">
      <router-view v-slot="{ Component }">
        <transition name="fade" mode="out-in">
          <component :is="Component" />
        </transition>
      </router-view>
    </main>
    <footer class="py-6 text-center text-xs text-slate-500 border-t border-white/5 bg-slate-950/30">
      <p>© 2026 AWS AI Academy & Dicoding Indonesia. Built with Vue 3 & Go Gin.</p>
    </footer>
  </div>
</template>
<style>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.25s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>

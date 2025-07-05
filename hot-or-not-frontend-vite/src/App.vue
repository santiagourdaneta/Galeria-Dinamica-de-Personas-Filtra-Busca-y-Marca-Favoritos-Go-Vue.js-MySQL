<script setup>
import PersonaGallery from './components/PersonaGallery.vue'; // Import the new component
</script>

<template>
  <div class="container">
    <header>
      <h1>Galeria de Personas Go+Vue+MySQL</h1>
      </header>
    <main>
      <PersonaGallery /> </main>
  </div>
</template>

<style scoped>
/* Estilos adicionales si los necesitas, aparte de Pico.css */
.filter-buttons {
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-bottom: 2rem;
}

.loading-state, .no-personas-state {
  text-align: center;
  padding: 2rem;
  border: 1px solid var(--pico-border-color);
  border-radius: var(--pico-border-radius);
}

article {
  max-width: 500px; /* Limita el ancho de la tarjeta para mejor UI */
  margin: auto;
  text-align: center;
}
</style>
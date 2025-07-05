<template>
  <div class="persona-card">
    <img :src="persona.url_imagen" :alt="persona.nombre" class="persona-image" />
    <hgroup>
      <h3>{{ persona.nombre }}</h3>
      <p>{{ persona.edad }} años</p>
    </hgroup>

    <div role="group" class="interaction-buttons">
      <button @click="$emit('dislike')" class="contrast outline">
        <i class="fas fa-times"></i> </button>
      <button @click="$emit('like')" class="secondary">
        <i class="fas fa-heart"></i> </button>
    </div>
  </div>
</template>

<script setup>


defineProps({
  persona: {
    type: Object,
    required: true,
  },
});

defineEmits(['like', 'dislike']);
</script>

<style scoped>
.persona-card {
  text-align: center;
  margin-top: 2rem;
  border: 1px solid var(--pico-border-color);
  border-radius: var(--pico-border-radius);
  padding: 1rem;
  box-shadow: var(--pico-card-box-shadow);
}

.persona-image {
  max-width: 100%;
  height: 400px; /* Altura fija para uniformidad */
  object-fit: cover; /* Recorta la imagen para llenar el espacio */
  border-radius: var(--pico-border-radius);
  margin-bottom: 1rem;
}

.interaction-buttons {
  display: flex;
  justify-content: center;
  gap: 1.5rem; /* Espacio entre los botones */
  margin-top: 1.5rem;
}

.interaction-buttons button {
  width: 60px; /* Tamaño de botón más grande */
  height: 60px;
  border-radius: 50%; /* Hacerlos circulares */
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 1.5rem; /* Iconos más grandes */
  transition: transform 0.2s ease-in-out, background-color 0.2s ease-in-out;
}

.interaction-buttons button:hover {
  transform: scale(1.1);
}

/* Colores para los botones */
.interaction-buttons button.secondary {
  background-color: var(--pico-color-red-500); /* Corazón rojo */
  border-color: var(--pico-color-red-500);
  color: var(--pico-color-white);
}

.interaction-buttons button.secondary:hover {
  background-color: var(--pico-color-red-600);
}

.interaction-buttons button.contrast.outline {
  background-color: var(--pico-color-gray-200); /* X gris claro */
  border-color: var(--pico-color-gray-200);
  color: var(--pico-color-gray-800);
}

.interaction-buttons button.contrast.outline:hover {
  background-color: var(--pico-color-gray-300);
}
</style>
<script setup>
import { ref, onMounted, watch } from 'vue';
import { getAllPersonas, toggleLike } from '../api/personaService'; // Import toggleLike

const personas = ref([]);
const currentFilter = ref('');
const searchQuery = ref('');

const fetchAllPersonas = async () => {
  try {
    const data = await getAllPersonas(currentFilter.value, searchQuery.value);
    if (data && Array.isArray(data)) {
      personas.value = data;
    } else {
      console.log("No personas found for the gallery with current filters.");
      personas.value = [];
    }
  } catch (error) {
    console.error("Error fetching personas for the gallery:", error);
    personas.value = [];
  }
};

const handleToggleLike = async (personaId) => {
  try {
    const result = await toggleLike(personaId);
    console.log(result.message);

    // Optimistically update the UI without refetching all personas
    // Find the persona in the current array and update its `is_liked` status
    const index = personas.value.findIndex(p => p.id === personaId);
    if (index !== -1) {
      personas.value[index].is_liked = result.is_liked;
    }
  } catch (error) {
    console.error("Error toggling like:", error);
    // You might want to revert UI changes or show an error message to the user
  }
};

onMounted(fetchAllPersonas);
watch([currentFilter, searchQuery], () => {
  fetchAllPersonas();
});

const resetFilters = () => {
  currentFilter.value = '';
  searchQuery.value = '';
};
</script>

<template>
  <div class="gallery-container">


    <div class="controls-bar">
      <select v-model="currentFilter" class="filter-select" aria-label="Filter by Gender">
        <option value="">All Genders</option>
        <option value="FEMENINO">Femenino</option>
        <option value="MASCULINO">Masculino</option>
      </select>
      <input
        type="text"
        v-model="searchQuery"
        placeholder="Search by name..."
        class="search-input"
        aria-label="Search by Name"
      />
      <button @click="resetFilters" class="reset-button">Reset Filters</button>
    </div>

    <div v-if="personas.length > 0" class="persona-grid">
      <div v-for="persona in personas" :key="persona.id" class="persona-item">
        <img
          :src="persona.url_imagen"
          :alt="`Portrait of ${persona.nombre}`"
          class="persona-image"
          loading="lazy"
        />
        <div class="persona-info">
          <h3>{{ persona.nombre }}, {{ persona.edad }}</h3>
          <p>{{ persona.genero }}</p>
          <button
            @click="handleToggleLike(persona.id)"
            :class="['like-button', { 'liked': persona.is_liked }]"
            :aria-label="persona.is_liked ? 'Unlike' : 'Like'"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              fill="currentColor"
              class="heart-icon"
            >
              <path
                d="M12 21.35l-1.84-1.83C5.46 15.52 2 12.02 2 8.5C2 5.42 4.42 3 7.5 3c1.74 0 3.41.81 4.5 2.09C13.09 3.81 14.76 3 16.5 3 19.58 3 22 5.42 22 8.5c0 3.52-3.46 7.02-8.16 10.95L12 21.35z"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>
    <div v-else class="no-results-message">
      <p>No personas found matching your criteria. Try adjusting the filters!</p>
    </div>
  </div>
</template>

<style scoped>
/* Keep existing .gallery-container, .filter-select, .search-input, .reset-button */

.controls-bar {
  display: flex;
  gap: 15px;
  margin-bottom: 25px; /* <--- AUMENTA ESTE VALOR PARA MÁS MARGEN */
  flex-wrap: wrap;
  align-items: center;
  padding-bottom: 15px; /* Opcional: añade un poco de padding interno si el espacio del margen no es suficiente */
  border-bottom: 1px solid #eee; /* Opcional: una línea sutil para separar */
}

.persona-grid {
  display: grid;
  /* Key change: Adjusted minmax value for more columns */
  /* This means each card will be at least 180px wide, and will grow to fill space */
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 15px; /* Slightly reduced gap for more compactness */
}

.persona-item {
  border: 1px solid #eee;
  border-radius: 8px;
  overflow: hidden;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  display: flex;
  flex-direction: column;
  transition: transform 0.2s ease-in-out, box-shadow 0.2s ease-in-out;
}

.persona-item:hover {
  transform: translateY(-5px);
  box-shadow: 0 6px 12px rgba(0, 0, 0, 0.15);
}

.persona-image {
  width: 100%;
  /* Maintain a good aspect ratio that shows faces well, e.g., 1:1 or 4:3 */
  /* For `object-fit: cover;`, a fixed height is often best */
  height: 200px; /* Reduced height slightly to make cards less tall */
  object-fit: cover; /* Ensures images fill the space without distortion, cropping as needed */
  display: block;
  background-color: #f0f0f0; /* Placeholder background */
}

.persona-info {
  padding: 12px; /* Slightly reduced padding */
  text-align: center;
  flex-grow: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.persona-info h3 {
  margin-top: 0;
  margin-bottom: 5px;
  color: #fff;
  font-size: 1.1em; /* Ensure text remains readable */
}

.persona-info p {
  color: #fff;
  font-size: 0.85em; /* Slightly smaller font for compactness */
  margin-bottom: 10px; /* Space before like button */
}

.like-button {
  background: none;
  border: none;
  cursor: pointer;
  padding: 0;
  width: 28px; /* Slightly smaller button */
  height: 28px;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0 auto;
  transition: transform 0.2s ease-in-out;
}

.like-button:hover {
  transform: scale(1.1);
}

.heart-icon {
  width: 100%;
  height: 100%;
  color: rgba(128, 128, 128, 0.5);
  fill: currentColor;
  transition: color 0.2s ease-in-out;
}

.like-button.liked .heart-icon {
  color: #ff0000;
}

.no-results-message {
  text-align: center;
  padding: 40px;
  font-size: 1.2em;
  color: #555;
}

/* Responsive adjustments for smaller screens (keep these or adjust based on new minmax) */
@media (max-width: 768px) {
  .controls-bar {
    flex-direction: column;
    gap: 10px;
  }
  .filter-select, .search-input, .reset-button {
    width: 100%;
  }
}
/* This media query can remain, as it will kick in when cards become too small */
@media (max-width: 600px) {
  .persona-grid {
    grid-template-columns: 1fr; /* Single column on very small screens */
  }
}
</style>
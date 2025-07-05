// Define API_BASE_URL at the top of the file
const API_BASE_URL = import.meta.env.VITE_APP_API_URL || 'http://localhost:8000/api';

export const getAllPersonas = async (genero = '', search = '') => {
    let url = `${API_BASE_URL}/personas`;
    const params = new URLSearchParams();

    if (genero) {
        params.append('genero', genero);
    }
    if (search) {
        params.append('search', search);
    }

    if (params.toString()) {
        url += `?${params.toString()}`;
    }

    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json();
};

export const toggleLike = async (personaId) => { // <--- ADD THIS NEW FUNCTION
    const response = await fetch(`${API_BASE_URL}/personas/${personaId}/toggle-like`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        // No body needed for this simple toggle, but good to include headers
    });
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
    return response.json(); // Returns { message: "...", is_liked: true/false }
};


export const likePersona = async (personaId) => {
    try {
        const response = await fetch(`${API_URL}/personas/${personaId}/like`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                // 'Authorization': `Bearer ${token}` // Si implementas JWT
            },
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error("Error liking persona:", error);
        throw error;
    }
};

export const dislikePersona = async (personaId) => {
    try {
        const response = await fetch(`${API_URL}/personas/${personaId}/dislike`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                // 'Authorization': `Bearer ${token}` // Si implementas JWT
            },
        });
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        return await response.json();
    } catch (error) {
        console.error("Error disliking persona:", error);
        throw error;
    }
};
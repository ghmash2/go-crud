// API base URL - change this if your backend runs on a different port
const API_BASE_URL = 'http://localhost:8080/api';

// Generic API call function
async function apiCall(endpoint, options = {}) {
    try {
        const response = await fetch(`${API_BASE_URL}${endpoint}`, {
            ...options,
            headers: {
                'Content-Type': 'application/json',
                ...options.headers,
            },
        });

        if (!response.ok) {
            const error = await response.text();
            throw new Error(error || `HTTP error! status: ${response.status}`);
        }

        // Return JSON if there's content
        const text = await response.text();
        return text ? JSON.parse(text) : null;
    } catch (error) {
        console.error('API Error:', error);
        throw error;
    }
}

// Users API
export const usersAPI = {
    getAll: () => apiCall('/users'),
    getById: (id) => apiCall(`/users/?id=${id}`),
    create: (user) => apiCall('/users', {
        method: 'POST',
        body: JSON.stringify(user),
    }),
    update: (id, user) => apiCall(`/users/?id=${id}`, {
        method: 'PUT',
        body: JSON.stringify(user),
    }),
    delete: (id) => apiCall(`/users/?id=${id}`, {
        method: 'DELETE',
    }),
};

// Products API
export const productsAPI = {
    getAll: () => apiCall('/products'),
    getById: (id) => apiCall(`/products/?id=${id}`),
    create: (product) => apiCall('/products', {
        method: 'POST',
        body: JSON.stringify(product),
    }),
    update: (id, product) => apiCall(`/products/?id=${id}`, {
        method: 'PUT',
        body: JSON.stringify(product),
    }),
    delete: (id) => apiCall(`/products/?id=${id}`, {
        method: 'DELETE',
    }),
};

// Categories API
export const categoriesAPI = {
    getAll: () => apiCall('/categories'),
    getById: (id) => apiCall(`/categories/?id=${id}`),
    create: (category) => apiCall('/categories', {
        method: 'POST',
        body: JSON.stringify(category),
    }),
    update: (id, category) => apiCall(`/categories/?id=${id}`, {
        method: 'PUT',
        body: JSON.stringify(category),
    }),
    delete: (id) => apiCall(`/categories/?id=${id}`, {
        method: 'DELETE',
    }),
};

// Utility functions
export function showMessage(message, type = 'success') {
    const messageDiv = document.createElement('div');
    messageDiv.className = `message message-${type}`;
    messageDiv.textContent = message;
    document.body.appendChild(messageDiv);

    setTimeout(() => {
        messageDiv.remove();
    }, 3000);
}

export function formatDate(dateString) {
    if (!dateString) return 'N/A';
    const date = new Date(dateString);
    return date.toLocaleDateString() + ' ' + date.toLocaleTimeString();
}

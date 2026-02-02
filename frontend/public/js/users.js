import { usersAPI, showMessage } from './api.js';

let isEditing = false;
let editingId = null;

// Load all users
async function loadUsers() {
    try {
        const users = await usersAPI.getAll();
        const tbody = document.getElementById('usersTableBody');

        if (!users || users.length === 0) {
            tbody.innerHTML = '<tr><td colspan="5" class="empty">No users found</td></tr>';
            return;
        }

        tbody.innerHTML = users.map(user => `
            <tr>
                <td>${user.ID || user.id}</td>
                <td>${user.Name || user.name}</td>
                <td>${user.Email || user.email}</td>
                <td>${new Date(user.CreatedAt || user.created_at).toLocaleDateString()}</td>
                <td class="actions">
                    <button onclick="editUser(${user.ID || user.id})" class="btn btn-sm btn-warning">Edit</button>
                    <button onclick="deleteUser(${user.ID || user.id})" class="btn btn-sm btn-danger">Delete</button>
                </td>
            </tr>
        `).join('');
    } catch (error) {
        showMessage('Error loading users: ' + error.message, 'error');
    }
}

// Show form
function showForm(editing = false) {
    isEditing = editing;
    document.getElementById('formSection').style.display = 'block';
    document.getElementById('formTitle').textContent = editing ? 'Edit User' : 'Add New User';

    if (!editing) {
        document.getElementById('userForm').reset();
        document.getElementById('userId').value = '';
    }
}

// Hide form
function hideForm() {
    document.getElementById('formSection').style.display = 'none';
    document.getElementById('userForm').reset();
    isEditing = false;
    editingId = null;
}

// Edit user
window.editUser = async function (id) {
    try {
        const user = await usersAPI.getById(id);
        document.getElementById('userId').value = user.ID || user.id;
        document.getElementById('name').value = user.Name || user.name;
        document.getElementById('email').value = user.Email || user.email;
        document.getElementById('password').value = ''; // Don't show password
        editingId = id;
        showForm(true);
    } catch (error) {
        showMessage('Error loading user: ' + error.message, 'error');
    }
};

// Delete user
window.deleteUser = async function (id) {
    if (!confirm('Are you sure you want to delete this user?')) return;

    try {
        await usersAPI.delete(id);
        showMessage('User deleted successfully');
        loadUsers();
    } catch (error) {
        showMessage('Error deleting user: ' + error.message, 'error');
    }
};

// Form submit
document.getElementById('userForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const userData = {
        name: document.getElementById('name').value,
        email: document.getElementById('email').value,
        password: document.getElementById('password').value,
    };

    try {
        if (isEditing) {
            await usersAPI.update(editingId, userData);
            showMessage('User updated successfully');
        } else {
            await usersAPI.create(userData);
            showMessage('User created successfully');
        }
        hideForm();
        loadUsers();
    } catch (error) {
        showMessage('Error saving user: ' + error.message, 'error');
    }
});

// Event listeners
document.getElementById('addBtn').addEventListener('click', () => showForm(false));
document.getElementById('cancelBtn').addEventListener('click', hideForm);

// Initial load
loadUsers();

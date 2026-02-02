import { categoriesAPI, showMessage } from './api.js';

let isEditing = false;
let editingId = null;

async function loadCategories() {
    try {
        const categories = await categoriesAPI.getAll();
        const tbody = document.getElementById('categoriesTableBody');

        if (!categories || categories.length === 0) {
            tbody.innerHTML = '<tr><td colspan="4" class="empty">No categories found</td></tr>';
            return;
        }

        tbody.innerHTML = categories.map(category => `
            <tr>
                <td>${category.ID || category.id}</td>
                <td>${category.Name || category.name}</td>
                <td>${new Date(category.CreatedAt || category.created_at).toLocaleDateString()}</td>
                <td class="actions">
                    <button onclick="editCategory(${category.ID || category.id})" class="btn btn-sm btn-warning">Edit</button>
                    <button onclick="deleteCategory(${category.ID || category.id})" class="btn btn-sm btn-danger">Delete</button>
                </td>
            </tr>
        `).join('');
    } catch (error) {
        showMessage('Error loading categories: ' + error.message, 'error');
    }
}

function showForm(editing = false) {
    isEditing = editing;
    document.getElementById('formSection').style.display = 'block';
    document.getElementById('formTitle').textContent = editing ? 'Edit Category' : 'Add New Category';

    if (!editing) {
        document.getElementById('categoryForm').reset();
        document.getElementById('categoryId').value = '';
    }
}

function hideForm() {
    document.getElementById('formSection').style.display = 'none';
    document.getElementById('categoryForm').reset();
    isEditing = false;
    editingId = null;
}

window.editCategory = async function (id) {
    try {
        const category = await categoriesAPI.getById(id);
        document.getElementById('categoryId').value = category.ID || category.id;
        document.getElementById('name').value = category.Name || category.name;
        editingId = id;
        showForm(true);
    } catch (error) {
        showMessage('Error loading category: ' + error.message, 'error');
    }
};

window.deleteCategory = async function (id) {
    if (!confirm('Are you sure you want to delete this category?')) return;

    try {
        await categoriesAPI.delete(id);
        showMessage('Category deleted successfully');
        loadCategories();
    } catch (error) {
        showMessage('Error deleting category: ' + error.message, 'error');
    }
};

document.getElementById('categoryForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const categoryData = {
        name: document.getElementById('name').value,
    };

    try {
        if (isEditing) {
            await categoriesAPI.update(editingId, categoryData);
            showMessage('Category updated successfully');
        } else {
            await categoriesAPI.create(categoryData);
            showMessage('Category created successfully');
        }
        hideForm();
        loadCategories();
    } catch (error) {
        showMessage('Error saving category: ' + error.message, 'error');
    }
});

document.getElementById('addBtn').addEventListener('click', () => showForm(false));
document.getElementById('cancelBtn').addEventListener('click', hideForm);

loadCategories();

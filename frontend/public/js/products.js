import { productsAPI, showMessage } from './api.js';

let isEditing = false;
let editingId = null;

async function loadProducts() {
    try {
        const products = await productsAPI.getAll();
        const tbody = document.getElementById('productsTableBody');

        if (!products || products.length === 0) {
            tbody.innerHTML = '<tr><td colspan="5" class="empty">No products found</td></tr>';
            return;
        }

        tbody.innerHTML = products.map(product => `
            <tr>
                <td>${product.ID || product.id}</td>
                <td>${product.Name || product.name}</td>
                <td>$${parseFloat(product.Price || product.price).toFixed(2)}</td>
                <td>${product.CategoryID || product.category_id}</td>
                <td class="actions">
                    <button onclick="editProduct(${product.ID || product.id})" class="btn btn-sm btn-warning">Edit</button>
                    <button onclick="deleteProduct(${product.ID || product.id})" class="btn btn-sm btn-danger">Delete</button>
                </td>
            </tr>
        `).join('');
    } catch (error) {
        showMessage('Error loading products: ' + error.message, 'error');
    }
}

function showForm(editing = false) {
    isEditing = editing;
    document.getElementById('formSection').style.display = 'block';
    document.getElementById('formTitle').textContent = editing ? 'Edit Product' : 'Add New Product';

    if (!editing) {
        document.getElementById('productForm').reset();
        document.getElementById('productId').value = '';
    }
}

function hideForm() {
    document.getElementById('formSection').style.display = 'none';
    document.getElementById('productForm').reset();
    isEditing = false;
    editingId = null;
}

window.editProduct = async function (id) {
    try {
        const product = await productsAPI.getById(id);
        document.getElementById('productId').value = product.ID || product.id;
        document.getElementById('name').value = product.Name || product.name;
        document.getElementById('price').value = product.Price || product.price;
        document.getElementById('category_id').value = product.CategoryID || product.category_id;
        editingId = id;
        showForm(true);
    } catch (error) {
        showMessage('Error loading product: ' + error.message, 'error');
    }
};

window.deleteProduct = async function (id) {
    if (!confirm('Are you sure you want to delete this product?')) return;

    try {
        await productsAPI.delete(id);
        showMessage('Product deleted successfully');
        loadProducts();
    } catch (error) {
        showMessage('Error deleting product: ' + error.message, 'error');
    }
};

document.getElementById('productForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const productData = {
        name: document.getElementById('name').value,
        price: document.getElementById('price').value,
        category_id: document.getElementById('category_id').value,
    };

    try {
        if (isEditing) {
            await productsAPI.update(editingId, productData);
            showMessage('Product updated successfully');
        } else {
            await productsAPI.create(productData);
            showMessage('Product created successfully');
        }
        hideForm();
        loadProducts();
    } catch (error) {
        showMessage('Error saving product: ' + error.message, 'error');
    }
});

document.getElementById('addBtn').addEventListener('click', () => showForm(false));
document.getElementById('cancelBtn').addEventListener('click', hideForm);

loadProducts();

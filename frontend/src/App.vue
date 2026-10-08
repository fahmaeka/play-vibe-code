<script setup>
import { ref, onMounted } from 'vue';
import { fetchHealth, getItems, createItem, deleteItem } from './services/api';

const health = ref(null);
const healthLoading = ref(false);
const healthError = ref(null);

const items = ref([]);
const itemsLoading = ref(false);
const itemsError = ref(null);

const newItem = ref({
  title: '',
  description: '',
  status: 'pending'
});
const isSubmitting = ref(false);

const notification = ref(null);

function showToast(message, type = 'info') {
  notification.value = { message, type };
  setTimeout(() => {
    notification.value = null;
  }, 4000);
}

async function loadHealth() {
  healthLoading.value = true;
  healthError.value = null;
  try {
    health.value = await fetchHealth();
  } catch (err) {
    healthError.value = err.message || 'Gagal terhubung ke Backend Echo';
    health.value = null;
  } finally {
    healthLoading.value = false;
  }
}

async function loadItems() {
  itemsLoading.value = true;
  itemsError.value = null;
  try {
    items.value = await getItems();
  } catch (err) {
    itemsError.value = err.message || 'Gagal memuat data dari database';
  } finally {
    itemsLoading.value = false;
  }
}

async function handleCreateItem() {
  if (!newItem.value.title.trim()) {
    showToast('Judul item wajib diisi!', 'warning');
    return;
  }

  isSubmitting.value = true;
  try {
    await createItem(newItem.value);
    showToast('Item berhasil ditambahkan ke MySQL!', 'success');
    newItem.value = { title: '', description: '', status: 'pending' };
    await loadItems();
  } catch (err) {
    showToast(err.message || 'Gagal menambahkan item', 'error');
  } finally {
    isSubmitting.value = false;
  }
}

async function handleDeleteItem(id) {
  try {
    await deleteItem(id);
    showToast('Item berhasil dihapus', 'info');
    await loadItems();
  } catch (err) {
    showToast(err.message || 'Gagal menghapus item', 'error');
  }
}

async function refreshAll() {
  await Promise.all([loadHealth(), loadItems()]);
}

onMounted(() => {
  refreshAll();
});
</script>

<template>
  <div class="container">
    <!-- Header -->
    <header class="header-card">
      <div class="brand">
        <div class="logo-badge">⚡</div>
        <div>
          <h1 class="title">Play Vibe Code</h1>
          <p class="subtitle">Fullstack Starter: Go (Echo) + MySQL + Vue 3 + Bun</p>
        </div>
      </div>
      <div class="tech-pills">
        <span class="pill pill-bun">Bun</span>
        <span class="pill pill-vue">Vue 3</span>
        <span class="pill pill-go">Go Echo</span>
        <span class="pill pill-mysql">MySQL</span>
      </div>
    </header>

    <!-- Toast Notification -->
    <transition name="fade">
      <div v-if="notification" :class="['toast', `toast-${notification.type}`]">
        {{ notification.message }}
      </div>
    </transition>

    <!-- Main Grid -->
    <main class="grid-layout">
      <!-- Status & Health Card -->
      <section class="card">
        <div class="card-header">
          <div class="card-title-wrap">
            <span class="icon">🛰️</span>
            <h2>System Health & Status</h2>
          </div>
          <button @click="refreshAll" :disabled="healthLoading" class="btn btn-secondary btn-sm">
            <span :class="{ 'spin': healthLoading }">🔄</span> Refresh
          </button>
        </div>

        <div class="status-grid">
          <!-- Backend Status -->
          <div class="status-box">
            <span class="status-label">Backend API (Go Echo)</span>
            <div class="status-val">
              <span v-if="health && health.status === 'ok'" class="badge badge-success">
                <span class="dot dot-success"></span> Online (Port 8080)
              </span>
              <span v-else class="badge badge-danger">
                <span class="dot dot-danger"></span> Offline / Disconnected
              </span>
            </div>
          </div>

          <!-- MySQL Status -->
          <div class="status-box">
            <span class="status-label">Database (MySQL)</span>
            <div class="status-val">
              <span v-if="health && health.database === 'connected'" class="badge badge-success">
                <span class="dot dot-success"></span> Connected
              </span>
              <span v-else class="badge badge-warning">
                <span class="dot dot-warning"></span> Disconnected
              </span>
            </div>
          </div>

          <!-- Environment -->
          <div class="status-box">
            <span class="status-label">Environment</span>
            <span class="text-highlight">{{ health?.environment || 'development' }}</span>
          </div>

          <!-- Uptime -->
          <div class="status-box">
            <span class="status-label">Uptime</span>
            <span class="text-highlight">{{ health?.uptime || '-' }}</span>
          </div>
        </div>

        <div v-if="healthError" class="alert alert-error mt-4">
          <strong>Perhatian:</strong> {{ healthError }}
          <div class="text-muted text-sm mt-1">Pastikan backend berjalan dengan perintah: <code>cd backend && go run ./cmd/api</code></div>
        </div>
      </section>

      <!-- CRUD Demo Card -->
      <section class="card">
        <div class="card-header">
          <div class="card-title-wrap">
            <span class="icon">📦</span>
            <h2>MySQL Data Store Demo</h2>
          </div>
          <span class="badge badge-neutral">{{ items.length }} Records</span>
        </div>

        <!-- Add Item Form -->
        <form @submit.prevent="handleCreateItem" class="form-grid">
          <div class="form-group">
            <label for="item-title">Title</label>
            <input 
              id="item-title" 
              v-model="newItem.title" 
              type="text" 
              placeholder="Contoh: Fitur Autentikasi" 
              class="input-control" 
              required
            />
          </div>

          <div class="form-group">
            <label for="item-desc">Deskripsi</label>
            <input 
              id="item-desc" 
              v-model="newItem.description" 
              type="text" 
              placeholder="Detail singkat tugas..." 
              class="input-control"
            />
          </div>

          <div class="form-group form-action">
            <button type="submit" :disabled="isSubmitting" class="btn btn-primary">
              <span>➕</span> {{ isSubmitting ? 'Menyimpan...' : 'Tambah Item' }}
            </button>
          </div>
        </form>

        <!-- Items Table / List -->
        <div class="items-container mt-4">
          <div v-if="itemsLoading" class="loading-state">
            <span class="spin">⏳</span> Memuat data dari MySQL...
          </div>

          <div v-else-if="itemsError" class="alert alert-warning">
            {{ itemsError }}
            <p class="text-sm mt-1">Nyalakan MySQL via Docker: <code>docker compose up -d</code> lalu restart backend.</p>
          </div>

          <div v-else-if="items.length === 0" class="empty-state">
            <span class="empty-icon">📭</span>
            <p>Belum ada data item di database MySQL.</p>
            <p class="text-sm text-muted">Gunakan form di atas untuk mencoba write data ke MySQL.</p>
          </div>

          <div v-else class="items-list">
            <div v-for="item in items" :key="item.id" class="item-card">
              <div class="item-info">
                <div class="item-header">
                  <span class="item-title">{{ item.title }}</span>
                  <span class="badge badge-subtle">#{{ item.id }}</span>
                </div>
                <p v-if="item.description" class="item-desc">{{ item.description }}</p>
                <div class="item-meta">
                  <span class="badge badge-status">{{ item.status }}</span>
                  <span class="timestamp">{{ new Date(item.created_at).toLocaleString() }}</span>
                </div>
              </div>
              <button 
                @click="handleDeleteItem(item.id)" 
                class="btn-icon btn-danger-icon" 
                title="Hapus Item"
              >
                🗑️
              </button>
            </div>
          </div>
        </div>
      </section>
    </main>

    <!-- Quick Start Guide Footer -->
    <footer class="footer-card">
      <h3>🚀 Quick Start Guide</h3>
      <div class="guide-grid">
        <div class="guide-item">
          <span class="step-num">1</span>
          <div>
            <strong>Nyalakan Database</strong>
            <code>docker compose up -d</code>
          </div>
        </div>
        <div class="guide-item">
          <span class="step-num">2</span>
          <div>
            <strong>Jalankan Backend</strong>
            <code>cd backend && go run ./cmd/api</code>
          </div>
        </div>
        <div class="guide-item">
          <span class="step-num">3</span>
          <div>
            <strong>Jalankan Frontend</strong>
            <code>cd frontend && bun dev</code>
          </div>
        </div>
      </div>
    </footer>
  </div>
</template>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  gap: 1.75rem;
}

.header-card {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem 2rem;
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-subtle);
  backdrop-filter: blur(12px);
  gap: 1rem;
}

.brand {
  display: flex;
  align-items: center;
  gap: 1.25rem;
}

.logo-badge {
  font-size: 2rem;
  width: 52px;
  height: 52px;
  background: rgba(56, 189, 248, 0.1);
  border: 1px solid var(--border-glow);
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: var(--radius-md);
}

.title {
  font-size: 1.6rem;
  font-weight: 700;
  background: var(--accent-gradient);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.subtitle {
  color: var(--text-secondary);
  font-size: 0.95rem;
}

.tech-pills {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}

.pill {
  padding: 0.35rem 0.85rem;
  border-radius: 9999px;
  font-size: 0.8rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.pill-bun { background: rgba(251, 146, 60, 0.15); color: #fb923c; border: 1px solid rgba(251, 146, 60, 0.3); }
.pill-vue { background: rgba(16, 185, 129, 0.15); color: #34d399; border: 1px solid rgba(16, 185, 129, 0.3); }
.pill-go { background: rgba(56, 189, 248, 0.15); color: #38bdf8; border: 1px solid rgba(56, 189, 248, 0.3); }
.pill-mysql { background: rgba(129, 140, 248, 0.15); color: #818cf8; border: 1px solid rgba(129, 140, 248, 0.3); }

.grid-layout {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1.75rem;
}

.card {
  background: var(--bg-card);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: 1.75rem;
  box-shadow: var(--shadow-subtle);
  backdrop-filter: blur(8px);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.25rem;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 0.85rem;
}

.card-title-wrap {
  display: flex;
  align-items: center;
  gap: 0.6rem;
}

.card-title-wrap h2 {
  font-size: 1.25rem;
  font-weight: 600;
}

.status-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 1rem;
}

.status-box {
  background: rgba(0, 0, 0, 0.25);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  padding: 1rem 1.25rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.status-label {
  font-size: 0.8rem;
  color: var(--text-muted);
  text-transform: uppercase;
  font-weight: 600;
}

.text-highlight {
  font-weight: 600;
  color: var(--text-primary);
}

.badge {
  display: inline-flex;
  align-items: center;
  gap: 0.4rem;
  padding: 0.25rem 0.75rem;
  border-radius: 9999px;
  font-size: 0.8rem;
  font-weight: 600;
}

.badge-success { background: rgba(16, 185, 129, 0.15); color: #34d399; }
.badge-danger { background: rgba(244, 63, 94, 0.15); color: #fb7185; }
.badge-warning { background: rgba(245, 158, 11, 0.15); color: #fbbf24; }
.badge-neutral { background: rgba(255, 255, 255, 0.1); color: var(--text-secondary); }
.badge-subtle { background: rgba(255, 255, 255, 0.05); color: var(--text-muted); }
.badge-status { background: rgba(56, 189, 248, 0.15); color: #38bdf8; text-transform: uppercase; font-size: 0.75rem; }

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  display: inline-block;
}
.dot-success { background-color: #10b981; box-shadow: 0 0 8px #10b981; }
.dot-danger { background-color: #f43f5e; box-shadow: 0 0 8px #f43f5e; }
.dot-warning { background-color: #f59e0b; box-shadow: 0 0 8px #f59e0b; }

.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr auto;
  gap: 1rem;
  align-items: flex-end;
}

@media (max-width: 768px) {
  .form-grid {
    grid-template-columns: 1fr;
  }
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.4rem;
}

.form-group label {
  font-size: 0.85rem;
  color: var(--text-secondary);
}

.input-control {
  background: rgba(0, 0, 0, 0.35);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-sm);
  padding: 0.65rem 0.9rem;
  color: var(--text-primary);
  font-size: 0.95rem;
  outline: none;
  transition: border-color 0.2s;
}

.input-control:focus {
  border-color: var(--accent-cyan);
}

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  padding: 0.65rem 1.25rem;
  border-radius: var(--radius-sm);
  font-weight: 600;
  font-size: 0.9rem;
  cursor: pointer;
  border: none;
  transition: all 0.2s;
}

.btn-primary {
  background: #0284c7;
  color: #ffffff;
}
.btn-primary:hover:not(:disabled) {
  background: #0369a1;
}

.btn-secondary {
  background: rgba(255, 255, 255, 0.08);
  color: var(--text-primary);
  border: 1px solid var(--border-color);
}
.btn-secondary:hover:not(:disabled) {
  background: rgba(255, 255, 255, 0.15);
}

.btn-sm {
  padding: 0.35rem 0.75rem;
  font-size: 0.8rem;
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}

.items-list {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.item-card {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem 1.25rem;
  background: rgba(0, 0, 0, 0.2);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-md);
  transition: background 0.2s;
}

.item-card:hover {
  background: var(--bg-card-hover);
}

.item-header {
  display: flex;
  align-items: center;
  gap: 0.6rem;
}

.item-title {
  font-weight: 600;
  font-size: 1.05rem;
}

.item-desc {
  color: var(--text-secondary);
  font-size: 0.9rem;
  margin-top: 0.25rem;
}

.item-meta {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-top: 0.5rem;
}

.timestamp {
  font-size: 0.75rem;
  color: var(--text-muted);
}

.btn-icon {
  background: none;
  border: none;
  cursor: pointer;
  font-size: 1.1rem;
  padding: 0.5rem;
  border-radius: var(--radius-sm);
  transition: background 0.2s;
}

.btn-danger-icon:hover {
  background: rgba(244, 63, 94, 0.15);
}

.empty-state, .loading-state {
  text-align: center;
  padding: 2.5rem 1rem;
  color: var(--text-secondary);
}

.empty-icon {
  font-size: 2.5rem;
  display: block;
  margin-bottom: 0.5rem;
}

.alert {
  padding: 0.85rem 1.15rem;
  border-radius: var(--radius-sm);
  font-size: 0.9rem;
}

.alert-error {
  background: rgba(244, 63, 94, 0.12);
  border: 1px solid rgba(244, 63, 94, 0.3);
  color: #fecdd3;
}

.alert-warning {
  background: rgba(245, 158, 11, 0.12);
  border: 1px solid rgba(245, 158, 11, 0.3);
  color: #fde68a;
}

.toast {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  padding: 0.85rem 1.5rem;
  border-radius: var(--radius-md);
  font-size: 0.9rem;
  font-weight: 500;
  box-shadow: 0 10px 25px rgba(0,0,0,0.5);
  z-index: 99;
}

.toast-success { background: #065f46; color: #a7f3d0; border: 1px solid #10b981; }
.toast-error { background: #881337; color: #fecdd3; border: 1px solid #f43f5e; }
.toast-info { background: #075985; color: #bae6fd; border: 1px solid #38bdf8; }
.toast-warning { background: #78350f; color: #fde68a; border: 1px solid #f59e0b; }

.footer-card {
  background: rgba(18, 24, 38, 0.6);
  border: 1px solid var(--border-color);
  border-radius: var(--radius-lg);
  padding: 1.5rem 2rem;
}

.footer-card h3 {
  font-size: 1.1rem;
  margin-bottom: 1rem;
}

.guide-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 1.25rem;
}

.guide-item {
  display: flex;
  align-items: flex-start;
  gap: 0.75rem;
}

.step-num {
  background: var(--accent-gradient);
  color: #000;
  font-weight: 700;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.8rem;
  flex-shrink: 0;
}

.guide-item code {
  display: block;
  background: rgba(0,0,0,0.4);
  padding: 0.35rem 0.6rem;
  border-radius: 4px;
  font-size: 0.85rem;
  color: #38bdf8;
  margin-top: 0.3rem;
}

.mt-1 { margin-top: 0.25rem; }
.mt-4 { margin-top: 1rem; }
.text-sm { font-size: 0.85rem; }
.spin {
  display: inline-block;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>

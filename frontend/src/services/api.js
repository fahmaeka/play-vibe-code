const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

export async function fetchHealth() {
  const response = await fetch(`${API_BASE_URL}/health`, {
    headers: { 'Accept': 'application/json' }
  });
  if (!response.ok) {
    throw new Error(`Health check failed with status: ${response.status}`);
  }
  return response.json();
}

export async function getItems() {
  const response = await fetch(`${API_BASE_URL}/api/items`, {
    headers: { 'Accept': 'application/json' }
  });
  if (!response.ok) {
    throw new Error(`Failed to fetch items: ${response.status}`);
  }
  return response.json();
}

export async function createItem(item) {
  const response = await fetch(`${API_BASE_URL}/api/items`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Accept': 'application/json',
    },
    body: JSON.stringify(item),
  });
  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}));
    throw new Error(errorData.error || `Failed to create item (${response.status})`);
  }
  return response.json();
}

export async function deleteItem(id) {
  const response = await fetch(`${API_BASE_URL}/api/items/${id}`, {
    method: 'DELETE',
    headers: { 'Accept': 'application/json' }
  });
  if (!response.ok) {
    throw new Error(`Failed to delete item (${response.status})`);
  }
  return response.json();
}

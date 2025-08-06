// src/app/categories/page.js

async function getCategories() {
  const res = await fetch('http://localhost:8080/categories');
  if (!res.ok) {
    throw new Error('Falha ao carregar categorias');
  }
  return res.json();
}

export default async function CategoriesPage() {
  const categories = await getCategories();

  return (
    <div>
      <h1>Categorias do Cardápio</h1>
      <ul>
        {categories.map((cat) => (
          <li key={cat.id}>{cat.name}</li>
        ))}
      </ul>
    </div>
  );
}

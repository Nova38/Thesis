<script lang="ts" setup>
const x = {
  id: '1000',
  code: 'f230fh0g3',
  name: 'Bamboo Watch',
  description: 'Product Description',
  image: 'bamboo-watch.jpg',
  price: 65,
  category: 'Accessories',
  quantity: 24,
  inventoryStatus: 'INSTOCK',
}

const products = ref([x, x, x])

for (let i = 0; i < 100; i++)
  products.value.push(x)

const columns = [
  { field: 'code', header: 'Code' },
  { field: 'name', header: 'Name' },
  { field: 'category', header: 'Category' },
  { field: 'quantity', header: 'Quantity' },
]
</script>

<template>
  <div>
    <PDataTable paginator :rows="5" :rows-per-page-options="[5, 10, 20, 50]" :value="products" table-style="min-width: 50rem">
      <PColumn v-for="col of columns" :key="col.field" :field="col.field" :header="col.header" />
    </PDataTable>
  </div>
</template>

<style scoped>

</style>

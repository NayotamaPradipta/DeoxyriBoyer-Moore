<template>
    <header class="flex items-center">
    <div style="height: 54px" class="pr-5">
        <img src="/logo-landscape.png" alt="logo" class="h-full" />
    </div>
    <ul class="flex items-center">
        <li>
        <nuxt-link
            class="text-white hover:text-teal-500 text-lg px-4 py-3"
            to="/"
            >Home</nuxt-link>
        </li>
        <li>
        <nuxt-link
            class="text-white hover:text-teal-500 text-lg px-4 py-3"
            to="/TambahPenyakit"
            >Tambah Penyakit</nuxt-link>
        </li>
        <li>
        <nuxt-link
            class="text-white hover:text-teal-500 text-lg px-4 py-3"
            to="/TesDNA"
            >Tes DNA</nuxt-link>
        </li>
        <li>
        <nuxt-link
            class="text-white hover:text-teal-500 text-lg px-4 py-3"
            to="/CariPenyakit"
            >Cari Penyakit</nuxt-link>
        </li>
    </ul>
    </header>
</template>
<template>
    <div class="h-screen flex justify-center items-center">
    <div
        class="hidden md:block lg:w-1/3 bg-white h-full auth-background rounded-tr-lg rounded-br-lg"
    ></div>
    <div class="w-auto md:w-2/4 lg:w-2/3 flex justify-center items-center">
        <div class="w-full lg:w-1/2 px-10 lg:px-0">
        <h2 class="font-normal mb-6 text-3xl text-white">
            Tambah Penyakit
        </h2>
        <div class="mb-6">
            <div class="mb-4">
            <label class="font-normal text-lg text-white block mb-3"
                >Nama Penyakit</label
            >
            <input
                type="text"
                class="auth-form focus:outline-none focus:bg-purple-hover focus:shadow-outline focus:border-purple-hover-stroke focus:text-gray-100"
                placeholder="Tulis penyakit di sini"
            />
            </div>
        </div>
        <div class="mb-6">
            <div class="mb-4">
            <label class="font-normal text-lg text-white block mb-3"
                >Sequence DNA</label
            >
            <input
                type="text"
                class="auth-form focus:outline-none focus:bg-purple-hover focus:shadow-outline focus:border-purple-hover-stroke focus:text-gray-100"
                placeholder="Unggah sequence DNA di sini"
            />
            </div>
        </div>
        <div class="mb-6">
            <div class="mb-4">
            <button
                @click="$router.push({ path: '/upload' })"
                class="block w-full bg-orange-button hover:bg-green-button text-white font-semibold px-6 py-4 text-lg rounded-full"
            >
                Submit
            </button>
            </div>
        </div>
        </div>
    </div>
    </div>
</template>

<script>
export default {
    layout: 'auth'
}
</script>

<style scoped>
    .auth-background {
    background-image: url("/sign-up-background.jpg");
    background-position: center;
    background-size: cover;
    }
</style>
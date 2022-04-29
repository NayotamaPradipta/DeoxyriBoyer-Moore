<template>
    <div class="landing-page">
      <section class="landing-hero pt-5">
        <div class="header__bg"></div>
        <div class="container mx-auto relative">
          <Navbar />
          <div class="flex items-center pt-10 px-5 md:px-0">
            <div class="w-1/2">
              <h1 class="text-4xl text-white mb-5">
                <u class="hero-underline">Tambah penyakit</u>, 
                <u class="hero-underline">tes DNA</u>, <br />
                dan <u class="hero-underline">cari penyakit</u>, AKURAT!* <br />
              </h1>
              <p class="text-white text-xl font-light mb-8">
                *kayaknya sih
              </p>
              <button
                @click="$router.push({ path: '/TesDNA' })"
                class="block bg-orange-button hover:bg-green-button text-white font-semibold px-12 py-3 text-xl rounded-full"
              >
                Tes DNA Sekarang!
              </button>
            </div>
            <div class="w-1/2 flex justify-center">
              <img src="/hero.png" alt="searching-dna" width="450" height="450"/>
            </div>
          </div>
        </div>
      </section>
      <Footer />
    </div>
</template>
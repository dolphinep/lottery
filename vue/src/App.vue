<template>
  <div style="backgroundColor: #2E2954; height: 100vh;">
   <div id="navbar">
        <nav class="navbar navbar-expand-lg navbar-light" style="backgroundColor:#B22335;">
            <a class="navbar-brand" style="color: #FFFFFF;">LOTTERY</a>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav">
                <li class="nav-item">
                    <a class="nav-link" href="/lottery" style="color: #FFFFFF; marginLeft: 10px;">Lotteries</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="/ticket" style="color: #FFFFFF; marginLeft: 10px;">Tickets</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="/announce" style="color: #FFFFFF; marginLeft: 10px;">Prize Announcements</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="/mywallet" style="color: #FFFFFF; marginLeft: 10px;">My Wallet</a>
                </li>
                </ul>
            </div>
            <sp-bank-balance-nav />
            <sp-sign-in />
            
        </nav>
    </div>
    <router-view />
  </div>
</template>

<style>
.sp-container {
  margin: 0 auto;
  max-width: 800px;
  padding: 1rem;
  background-color: #2E2954;
}
</style>

<script>
import * as sp from "@tendermint/vue";

export default {
  created() {
    this.$store.dispatch("cosmos/init");
  },
  components: { ...sp },
};
</script>

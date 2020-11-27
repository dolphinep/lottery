<template>
  <div style="backgroundColor: #2E2954;">
    <div class="sp-container">
      <sp-bank-balances />
      <sp-token-send />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
    name:  'MyWallet',
  components: { ...sp },
};
</script>
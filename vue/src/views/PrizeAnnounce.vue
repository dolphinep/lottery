<template>
  <div style="backgroundColor: #2E2954;">
    

    <div class="sp-container">
      
        <sp-type-form type="prizeAnnounce" :fields="['lotteryID', ]" module="lottery" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  name:  'PrizeAnnounce',
  components: { ...sp },
};
</script>

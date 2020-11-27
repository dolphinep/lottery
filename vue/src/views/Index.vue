<template>
  <div style="backgroundColor: #2E2954;">
    

    <div class="sp-container">
      
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form type="lottery" :fields="['name', 'detail', 'reward', 'drawdate', 'price', ]" module="lottery" />
		<sp-type-form type="ticket" :fields="['lotteryID', 'number', ]" module="lottery" />
		<sp-type-form type="prizeAnnounce" :fields="['lotteryID', ]" module="lottery" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>

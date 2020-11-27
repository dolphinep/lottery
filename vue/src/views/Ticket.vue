<template>
  <div style="backgroundColor: #2E2954;">
    
    <div class="sp-container">
      	<sp-type-form type="ticket" :fields="['lotteryID', 'number', ]" module="lottery" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  name:  'Ticket',
  components: { ...sp },
};
</script>

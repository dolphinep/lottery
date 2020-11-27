import Vue from "vue";
import VueRouter from "vue-router";
import Index from "../views/Index.vue";
import PrizeAnnounce from '../views/PrizeAnnounce.vue';
import Ticket from '../views/Ticket.vue';
import Lottery from '../views/Lottery.vue';
import MyWallet from '../views/MyWallet.vue';
Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    component: Lottery,
  },
//   { path: '/signin', component: SignIn },
//   { path: '/lotterylist', component: LotteryList },
//   { path: '/ticketlist',component: TicketList },
  { path: '/lottery', component: Lottery },
  { path: '/ticket',component: Ticket },
  { path: '/announce', component: PrizeAnnounce },
  { path: '/mywallet',component: MyWallet }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;

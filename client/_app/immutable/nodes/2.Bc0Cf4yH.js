import{s as A,B as D,a as d,e as E,g as S,c as F,i as _,D as L,E as M,F as P,f as b,k as j,o as z}from"../chunks/scheduler.Cdj48xKw.js";import{S as G,i as H,d as g,e as k,m as T,a as c,t as u,f as w}from"../chunks/index.krINOmAy.js";import{g as J}from"../chunks/entry.Cwgp6uEs.js";import{T as K,a as v}from"../chunks/Tabs.Cg4KEYbt.js";import{d as N,w as O}from"../chunks/index.DGw2-XJO.js";import{p as Q}from"../chunks/stores.B-q5h_hO.js";const U=()=>{const{subscribe:a,set:s}=O("teams");return{subscribe:a,set:r=>{s(r),localStorage.setItem("activeTab",r)},init:()=>{{const r=localStorage.getItem("activeTab");r&&s(r)}}}},C=U(),V=N(Q,a=>{const r=a.url.pathname.split("/").pop();return["teams","pools","seeding","bracket","results"].includes(r)?(C.set(r),r):"teams"});function W(a){let s,r,l,f,m,$,n,t,i,p;return s=new v({props:{open:a[0]==="teams",title:"Teams"}}),s.$on("click",a[3]),l=new v({props:{open:a[0]==="pools",title:"Pools"}}),l.$on("click",a[4]),m=new v({props:{open:a[0]==="seeding",title:"Seeding"}}),m.$on("click",a[5]),n=new v({props:{open:a[0]==="bracket",title:"Bracket"}}),n.$on("click",a[6]),i=new v({props:{open:a[0]==="results",title:"Results"}}),i.$on("click",a[7]),{c(){g(s.$$.fragment),r=d(),g(l.$$.fragment),f=d(),g(m.$$.fragment),$=d(),g(n.$$.fragment),t=d(),g(i.$$.fragment)},l(e){k(s.$$.fragment,e),r=S(e),k(l.$$.fragment,e),f=S(e),k(m.$$.fragment,e),$=S(e),k(n.$$.fragment,e),t=S(e),k(i.$$.fragment,e)},m(e,o){T(s,e,o),_(e,r,o),T(l,e,o),_(e,f,o),T(m,e,o),_(e,$,o),T(n,e,o),_(e,t,o),T(i,e,o),p=!0},p(e,o){const h={};o&1&&(h.open=e[0]==="teams"),s.$set(h);const B={};o&1&&(B.open=e[0]==="pools"),l.$set(B);const I={};o&1&&(I.open=e[0]==="seeding"),m.$set(I);const R={};o&1&&(R.open=e[0]==="bracket"),n.$set(R);const q={};o&1&&(q.open=e[0]==="results"),i.$set(q)},i(e){p||(c(s.$$.fragment,e),c(l.$$.fragment,e),c(m.$$.fragment,e),c(n.$$.fragment,e),c(i.$$.fragment,e),p=!0)},o(e){u(s.$$.fragment,e),u(l.$$.fragment,e),u(m.$$.fragment,e),u(n.$$.fragment,e),u(i.$$.fragment,e),p=!1},d(e){e&&(b(r),b(f),b($),b(t)),w(s,e),w(l,e),w(m,e),w(n,e),w(i,e)}}}function X(a){let s,r,l,f,m;s=new K({props:{class:"font-heading ml-2",contentClass:"",activeClasses:"p-4 text-theme border-b-2 border-theme",$$slots:{default:[W]},$$scope:{ctx:a}}});const $=a[2].default,n=D($,a,a[8],null);return{c(){g(s.$$.fragment),r=d(),l=E("br"),f=d(),n&&n.c()},l(t){k(s.$$.fragment,t),r=S(t),l=F(t,"BR",{}),f=S(t),n&&n.l(t)},m(t,i){T(s,t,i),_(t,r,i),_(t,l,i),_(t,f,i),n&&n.m(t,i),m=!0},p(t,[i]){const p={};i&257&&(p.$$scope={dirty:i,ctx:t}),s.$set(p),n&&n.p&&(!m||i&256)&&L(n,$,t,t[8],m?P($,t[8],i,null):M(t[8]),null)},i(t){m||(c(s.$$.fragment,t),c(n,t),m=!0)},o(t){u(s.$$.fragment,t),u(n,t),m=!1},d(t){t&&(b(r),b(l),b(f)),w(s,t),n&&n.d(t)}}}function Y(a,s,r){let l;j(a,V,o=>r(0,l=o));let{$$slots:f={},$$scope:m}=s;z(()=>{C.init()});function $(o){J(`./${o}`),C.set(o)}const n=()=>$("teams"),t=()=>$("pools"),i=()=>$("seeding"),p=()=>$("bracket"),e=()=>$("results");return a.$$set=o=>{"$$scope"in o&&r(8,m=o.$$scope)},[l,$,f,n,t,i,p,e,m]}class ne extends G{constructor(s){super(),H(this,s,Y,X,A,{})}}export{ne as component};

import{s as j,l as h,i as g,f as $,a as x,g as C,t as k,d as y,O as H,e as N,c as O,b as P,x as z,h as I,j as R}from"../chunks/scheduler.Cdj48xKw.js";import{S as q,i as A,t as i,c as B,a as f,g as S,d as m,e as _,m as p,f as d}from"../chunks/index.krINOmAy.js";import{T as D,a as E,b as L,c as M,e as v,d as U,f as V}from"../chunks/TableHeadCell.DCsmDhEb.js";import{H as F}from"../chunks/Heading.BAZVKAx1.js";const G="https://awcu2nks23.us-east-1.awsapprunner.com",J=async({params:o})=>{const e=await fetch(`${G}/api/${o.event}/teams`);return e.status==404?{teams:null}:{teams:await e.json()}},fe=Object.freeze(Object.defineProperty({__proto__:null,load:J},Symbol.toStringTag,{value:"Module"}));function w(o,e,l){const t=o.slice();return t[1]=e[l],t}function K(o){let e,l;return e=new D({props:{divClass:"ml-2 mr-2 font-default",$$slots:{default:[te]},$$scope:{ctx:o}}}),{c(){m(e.$$.fragment)},l(t){_(e.$$.fragment,t)},m(t,a){p(e,t,a),l=!0},p(t,a){const n={};a&17&&(n.$$scope={dirty:a,ctx:t}),e.$set(n)},i(t){l||(f(e.$$.fragment,t),l=!0)},o(t){i(e.$$.fragment,t),l=!1},d(t){d(e,t)}}}function Q(o){let e,l;return e=new F({props:{tag:"h5",class:"font-heading ml-2",customSize:"text-xl",$$slots:{default:[ae]},$$scope:{ctx:o}}}),{c(){m(e.$$.fragment)},l(t){_(e.$$.fragment,t)},m(t,a){p(e,t,a),l=!0},p(t,a){const n={};a&16&&(n.$$scope={dirty:a,ctx:t}),e.$set(n)},i(t){l||(f(e.$$.fragment,t),l=!0)},o(t){i(e.$$.fragment,t),l=!1},d(t){d(e,t)}}}function W(o){let e;return{c(){e=k("Name")},l(l){e=y(l,"Name")},m(l,t){g(l,e,t)},d(l){l&&$(e)}}}function X(o){let e,l;return e=new M({props:{$$slots:{default:[W]},$$scope:{ctx:o}}}),{c(){m(e.$$.fragment)},l(t){_(e.$$.fragment,t)},m(t,a){p(e,t,a),l=!0},p(t,a){const n={};a&16&&(n.$$scope={dirty:a,ctx:t}),e.$set(n)},i(t){l||(f(e.$$.fragment,t),l=!0)},o(t){i(e.$$.fragment,t),l=!1},d(t){d(e,t)}}}function Y(o){let e,l=o[1].name+"",t;return{c(){e=N("div"),t=k(l),this.h()},l(a){e=O(a,"DIV",{class:!0});var n=P(e);t=y(n,l),n.forEach($),this.h()},h(){z(e,"class","text-black")},m(a,n){g(a,e,n),I(e,t)},p(a,n){n&1&&l!==(l=a[1].name+"")&&R(t,l)},d(a){a&&$(e)}}}function Z(o){let e,l,t;return e=new V({props:{class:"py-2",$$slots:{default:[Y]},$$scope:{ctx:o}}}),{c(){m(e.$$.fragment),l=x()},l(a){_(e.$$.fragment,a),l=C(a)},m(a,n){p(e,a,n),g(a,l,n),t=!0},p(a,n){const s={};n&17&&(s.$$scope={dirty:n,ctx:a}),e.$set(s)},i(a){t||(f(e.$$.fragment,a),t=!0)},o(a){i(e.$$.fragment,a),t=!1},d(a){a&&$(l),d(e,a)}}}function T(o){let e,l;return e=new U({props:{color:"default",$$slots:{default:[Z]},$$scope:{ctx:o}}}),{c(){m(e.$$.fragment)},l(t){_(e.$$.fragment,t)},m(t,a){p(e,t,a),l=!0},p(t,a){const n={};a&17&&(n.$$scope={dirty:a,ctx:t}),e.$set(n)},i(t){l||(f(e.$$.fragment,t),l=!0)},o(t){i(e.$$.fragment,t),l=!1},d(t){d(e,t)}}}function ee(o){let e,l,t=v(o[0].teams),a=[];for(let s=0;s<t.length;s+=1)a[s]=T(w(o,t,s));const n=s=>i(a[s],1,1,()=>{a[s]=null});return{c(){for(let s=0;s<a.length;s+=1)a[s].c();e=h()},l(s){for(let c=0;c<a.length;c+=1)a[c].l(s);e=h()},m(s,c){for(let r=0;r<a.length;r+=1)a[r]&&a[r].m(s,c);g(s,e,c),l=!0},p(s,c){if(c&1){t=v(s[0].teams);let r;for(r=0;r<t.length;r+=1){const u=w(s,t,r);a[r]?(a[r].p(u,c),f(a[r],1)):(a[r]=T(u),a[r].c(),f(a[r],1),a[r].m(e.parentNode,e))}for(S(),r=t.length;r<a.length;r+=1)n(r);B()}},i(s){if(!l){for(let c=0;c<t.length;c+=1)f(a[c]);l=!0}},o(s){a=a.filter(Boolean);for(let c=0;c<a.length;c+=1)i(a[c]);l=!1},d(s){s&&$(e),H(a,s)}}}function te(o){let e,l,t,a;return e=new E({props:{class:"bg-theme text-white",$$slots:{default:[X]},$$scope:{ctx:o}}}),t=new L({props:{tableBodyClass:"divide-y",$$slots:{default:[ee]},$$scope:{ctx:o}}}),{c(){m(e.$$.fragment),l=x(),m(t.$$.fragment)},l(n){_(e.$$.fragment,n),l=C(n),_(t.$$.fragment,n)},m(n,s){p(e,n,s),g(n,l,s),p(t,n,s),a=!0},p(n,s){const c={};s&16&&(c.$$scope={dirty:s,ctx:n}),e.$set(c);const r={};s&17&&(r.$$scope={dirty:s,ctx:n}),t.$set(r)},i(n){a||(f(e.$$.fragment,n),f(t.$$.fragment,n),a=!0)},o(n){i(e.$$.fragment,n),i(t.$$.fragment,n),a=!1},d(n){n&&$(l),d(e,n),d(t,n)}}}function ae(o){let e;return{c(){e=k("There are no teams registered for this division yet. Check back later!")},l(l){e=y(l,"There are no teams registered for this division yet. Check back later!")},m(l,t){g(l,e,t)},d(l){l&&$(e)}}}function le(o){let e,l,t,a;const n=[Q,K],s=[];function c(r,u){return r[0].teams===null?0:1}return e=c(o),l=s[e]=n[e](o),{c(){l.c(),t=h()},l(r){l.l(r),t=h()},m(r,u){s[e].m(r,u),g(r,t,u),a=!0},p(r,[u]){let b=e;e=c(r),e===b?s[e].p(r,u):(S(),i(s[b],1,1,()=>{s[b]=null}),B(),l=s[e],l?l.p(r,u):(l=s[e]=n[e](r),l.c()),f(l,1),l.m(t.parentNode,t))},i(r){a||(f(l),a=!0)},o(r){i(l),a=!1},d(r){r&&$(t),s[e].d(r)}}}function ne(o,e,l){let{data:t}=e;return o.$$set=a=>{"data"in a&&l(0,t=a.data)},[t]}class ie extends q{constructor(e){super(),A(this,e,ne,le,j,{data:0})}}export{ie as component,fe as universal};

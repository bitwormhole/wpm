(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-b3c0de3e"],{5949:function(e,t,n){},"86d9":function(e,t,n){"use strict";n("5949")},a071:function(e,t,n){"use strict";n("c3c4")},b711:function(e,t,n){"use strict";n.r(t);var c=n("7a23");function o(e,t,n,o,r,a){var l=Object(c["resolveComponent"])("MyChild");return Object(c["openBlock"])(),Object(c["createBlock"])(l)}function r(e,t,n,o,r,a){var l=Object(c["resolveComponent"])("MyApps"),i=Object(c["resolveComponent"])("HomeFrame");return Object(c["openBlock"])(),Object(c["createBlock"])(i,{config:a.theFrameConfig,onAction:a.handleAction},{default:Object(c["withCtx"])((function(){return[Object(c["createVNode"])(l,{value:a.theItems},null,8,["value"])]})),_:1},8,["config","onAction"])}n("b0c0");var a={class:"list"};function l(e,t,n,o,r,l){var i=Object(c["resolveComponent"])("MyItem");return Object(c["openBlock"])(),Object(c["createElementBlock"])("div",a,[(Object(c["openBlock"])(!0),Object(c["createElementBlock"])(c["Fragment"],null,Object(c["renderList"])(n.value,(function(e){return Object(c["openBlock"])(),Object(c["createElementBlock"])("div",{class:"item",key:e.id},[Object(c["createVNode"])(i,{value:e},null,8,["value"])])})),128))])}function i(e,t,n,o,r,a){var l=Object(c["resolveComponent"])("MyOverview"),i=Object(c["resolveComponent"])("MyDetail"),s=Object(c["resolveComponent"])("el-popover");return Object(c["openBlock"])(),Object(c["createBlock"])(s,{placement:"top-start",width:380,trigger:"hover",content:"this is content, this is content, this is content",onShow:a.handlePopoverShow},{reference:Object(c["withCtx"])((function(){return[Object(c["createVNode"])(l,{value:n.value},null,8,["value"])]})),default:Object(c["withCtx"])((function(){return[Object(c["createVNode"])(i,{value:n.value},null,8,["value"])]})),_:1},8,["onShow"])}n("a4d3"),n("e01a");var s={class:"description"},u={class:"path"},d=Object(c["createTextVNode"])(" 运行 ");function v(e,t,n,o,r,a){var l=Object(c["resolveComponent"])("el-divider"),i=Object(c["resolveComponent"])("el-button");return Object(c["openBlock"])(),Object(c["createElementBlock"])("div",null,[Object(c["createElementVNode"])("div",null,[Object(c["createElementVNode"])("h2",null,Object(c["toDisplayString"])(n.value.name),1)]),Object(c["createElementVNode"])("div",s,Object(c["toDisplayString"])(n.value.description),1),Object(c["createVNode"])(l),Object(c["createElementVNode"])("div",u,Object(c["toDisplayString"])(n.value.path),1),Object(c["createVNode"])(i,{type:"success",onClick:a.handleRun},{default:Object(c["withCtx"])((function(){return[d]})),_:1},8,["onClick"])])}var p={name:"home-exe-detail",methods:{handleRun:function(){var e=this.value,t={executable:e},n={intents:[t]};this.$store.dispatch("intent/run",n)}},props:{value:Object}},b=n("6b0d"),m=n.n(b);const h=m()(p,[["render",v]]);var O=h,j={class:"item"},f={class:"icon-box"},k=["src"],y={class:"title-and-description"},C={class:"title"},w={class:"description"};function x(e,t,n,o,r,a){return Object(c["openBlock"])(),Object(c["createElementBlock"])("div",j,[Object(c["createElementVNode"])("div",f,[Object(c["createElementVNode"])("img",{class:"icon-image",src:n.value.icon},null,8,k)]),Object(c["createElementVNode"])("div",y,[Object(c["createElementVNode"])("div",C,Object(c["toDisplayString"])(n.value.title),1),Object(c["createElementVNode"])("div",w,Object(c["toDisplayString"])(n.value.description),1)])])}var g={name:"home-exe-overview",props:{value:Object}};n("86d9");const B=m()(g,[["render",x],["__scopeId","data-v-92e3dbbc"]]);var N=B,V={name:"home-exe-item",components:{MyDetail:O,MyOverview:N},methods:{handlePopoverShow:function(){}},props:{value:Object}};const E=m()(V,[["render",i]]);var M=E,S={name:"home-exe-list",components:{MyItem:M},props:{value:Array}};n("a071");const F=m()(S,[["render",l],["__scopeId","data-v-1643f5c1"]]);var A=F,D={name:"home-executable-index",components:{MyApps:A},computed:{theFrameConfig:function(){return{pager:!1,home:!0,nav:!0,search:!0,searchCategory:"应用"}},theItems:function(){return this.$store.getters["executable/afterFilter"]}},data:function(){return{}},methods:{handleAction:function(e){var t=e.name;"refresh"==t?this.fetch():"add"==t?this.add():"applyFilter"==t?this.applyFilter(e.value):console.warn("unsupported action: "+t)},applyFilter:function(e){var t=e;this.$store.commit("executable/update",{filter:t})},fetch:function(){this.$store.dispatch("executable/fetch")},add:function(){this.$router.push("/import-executable")}},mounted:function(){this.fetch()}};const _=m()(D,[["render",r]]);var I=_,$={name:"HomeExecutables",components:{MyChild:I}};const H=m()($,[["render",o]]);t["default"]=H},c3c4:function(e,t,n){}}]);
//# sourceMappingURL=chunk-b3c0de3e.4de4dd80.js.map
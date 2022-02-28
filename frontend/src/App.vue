<template>
<div>
<edit/>
<delete_modal/>
<main_html/>
</div>
</template>
<script>
import main_html from './components/main_html.vue'
import delete_modal from './components/delete_modal.vue'
import edit from './components/edit.vue'
export default {
  name: 'App',
  components: {
    edit,
    delete_modal,
    main_html,
  }
}
</script>


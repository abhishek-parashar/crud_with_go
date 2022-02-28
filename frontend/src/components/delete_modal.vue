<template>
  <div>
  <!-- Modal Window  for delete request  -->
  <div id="delete" class="modal fade" role="dialog">
    <div class="modal-dialog">
      <!-- Modal content-->
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal">&times;</button>
          <h4 class="modal-title">Delete</h4>
        </div>
        <div class="modal-body">
          <p>Are You sure you want to delete</p>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-primary" id="delreq" v-on:click="deletereq()">Yes</button>
          <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
        </div>
      </div>
    </div>
  </div>
  </div>
</template>

<script>
export default {
  name:'delete_modal',
  methods:{
    deletereq() {
      let url = 'http://localhost:8080/users?ref=1'
      let params = {
        method: 'DELETE',
        headers: {
            'content-type': 'application/json',
        },
        body: {},
      };
      fetch(url, params).then(res => res.json()
      ).then(data => console.warn(data)
      ).catch(err => console.error(err));
    }
  }
}
</script>

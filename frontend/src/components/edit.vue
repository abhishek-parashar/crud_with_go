<template>
  <div>
  <!-- Modal Window  for edit page  -->
  <div id="edit" class="modal fade" role="dialog">
    <div class="modal-dialog">
      <!-- Modal content-->
      <div class="modal-content">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal">&times;</button>
          <h4 class="modal-title">Edit User</h4>
        </div>
        <div class="modal-body">
          <form>
            <div class="form-group row">
              <label for="Name" class="col-sm-2 col-form-label">Name</label>
              <div class="col-sm-10">
                <input type="text" class="form-control" id="name_form"  placeholder="Name" v-model="name_form">
              </div>
            </div>
            <div class="form-group row">
              <label for="Team" class="col-sm-2 col-form-label">Team Name</label>
              <div class="col-sm-10">
                <input type="text" class="form-control" id="team_form" placeholder="Team Name" v-model="team_form">
              </div>=
            </div>
            <div class="form-group row">
              <label for="Location" class="col-sm-2 col-form-label">Location Name</label>
              <div class="col-sm-10">
                <input type="text" class="form-control" id="location_form" placeholder="Location Name" v-model="location_form">
              </div>
            </div>
          </form>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-primary" v-on:click="putreq()" >Update</button>
          <button type="button" class="btn btn-primary" data-toggle="modal" data-target="#delete">Delete</button>
        </div>
      </div>
    </div>
  </div>
  </div>
  </template>
<script>
export default {
  name: 'edit',
  data(){
    return {
      name_form:null,
      team_form:null,
      location_form:null,
    }
  },
  methods: {
    putreq() {
    let url = 'http://localhost:8080/users?ref=1';
    let name_elem = this.name_form
    let team_elem = this.team_form
    let location_elem = this.location_form
    console.log(name_elem)
    let datax = { user_name: name_elem, team_name: team_elem, location_name: location_elem };
    let params = {
        method: 'PUT',
        headers: {
            'content-type': 'application/json',
        },
        body: JSON.stringify(datax),
      };
    fetch(url, params).then(res => res.json()
      ).then(data => console.log(data)
      ).catch(err => console.error(err));
    },
  },
}
</script>


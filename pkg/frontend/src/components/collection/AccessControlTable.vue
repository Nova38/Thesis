<template>
  <div class="q-pa-lg" id="componentWrapper">
    <q-card>
      <q-card-section class="bg-secondary">
        <h4>Collection Role Permissions</h4>
      </q-card-section>

      <!-- Users perms -->
      <q-card-section v-if="data.users">
        <div class="text-h5 text-center">Users Actions</div>
        <div class="text-caption text-center">
          These permissions are for changing the roles of users in the
          collection
        </div>

        <div class="wrapper" v-if="data?.users">
          <div v-for="(value, key) in data.users" :key="key">
            <div class="permSection">
              <div class="permName">{{ key }}</div>

              <div class="permOptions">
                <q-option-group
                  inline
                  :disable="!isEditable"
                  :modelValue="data.users[key]"
                  :options="RoleOptions"
                  color="green"
                  type="checkbox"
                  @update:modelValue="
                    data.users = { ...data.users, [key]: $event }
                  "
                />
              </div>
            </div>
          </div>
        </div>
      </q-card-section>

      <q-separator />
      <!-- Role perms -->

      <q-card-section v-if="data.roles">
        <div class="text-h5 text-center">Role Permissions</div>
        <div class="text-caption text-center">
          These permissions are for managing what the roles can do in the
          collection
        </div>

        <div v-for="(value, key) in data?.roles" :key="key" class="permSection">
          <div class="permName">{{ key }}</div>

          <q-option-group
            inline
            :disable="!isEditable"
            :modelValue="data?.roles[key]"
            :options="RoleOptions"
            color="green"
            type="checkbox"
            @update:modelValue="
              data.roles = {
                ...data.roles,
                [key]: $event,
              }
            "
          />
        </div>
      </q-card-section>

      <q-separator />
      <q-card-section v-if="data.specimen">
        <div class="text-h5 text-center">Specimen Actions</div>
        <div class="text-caption text-center">
          These permissions are for managing the specimens in the collection
        </div>

        <div
          v-for="(value, key) in data.specimen"
          :key="key"
          class="permSection"
        >
          <div class="permName">{{ key }}</div>

          <q-option-group
            inline
            :disable="!isEditable"
            :modelValue="data.specimen[key]"
            :options="RoleOptions"
            color="green"
            type="checkbox"
            @update:modelValue="
              data.specimen = { ...data.specimen, [key]: $event }
            "
          />
        </div>
      </q-card-section>

      <!-- hidden -->
      <q-separator />

      <!-- Suggested -->
      <q-separator />

      <!-- primary -->
      <q-separator />
      <q-card-section v-if="data.primary">
        <div class="text-h5 text-center">Primary Actions</div>
        <div class="text-caption text-center">
          These permissions are for managing the specimens primary data fields
        </div>

        <div
          v-for="(value, key) in data.primary"
          :key="key"
          class="permSection"
        >
          <div class="permName">{{ key }}</div>

          <q-option-group
            inline
            :disable="!isEditable"
            :modelValue="data.primary[key]"
            :options="RoleOptions"
            color="green"
            type="checkbox"
            @update:modelValue="
              data.primary = { ...data.primary, [key]: $event }
            "
          />
        </div>
      </q-card-section>

      <!-- secondary -->
      <q-separator />
      <q-card-section v-if="data.secondary">
        <div class="text-h5 text-center">Secondary Actions</div>
        <div class="text-caption text-center">
          These permissions are for managing the specimens secondary data fields
        </div>

        <div
          v-for="(value, key) in data.secondary"
          :key="key"
          class="permSection"
        >
          <div class="permName">{{ key }}</div>

          <q-option-group
            inline
            :disable="!isEditable"
            :modelValue="data.secondary[key]"
            :options="RoleOptions"
            color="green"
            type="checkbox"
            @update:modelValue="
              data.secondary = { ...data.secondary, [key]: $event }
            "
          />
        </div>
      </q-card-section>

      <!-- taxon -->
      <q-separator />
      <q-card-section v-if="data.taxon">
        <div class="text-h5 text-center">Taxon Actions</div>
        <div class="text-caption text-center">
          These permissions are for managing the specimens taxon data fields
        </div>

        <div v-for="(value, key) in data.taxon" :key="key" class="permSection">
          <div class="permName">{{ key }}</div>

          <q-option-group
            inline
            :disable="!isEditable"
            :modelValue="data.taxon[key]"
            :options="RoleOptions"
            color="green"
            type="checkbox"
            @update:modelValue="data.taxon = { ...data.taxon, [key]: $event }"
          />
        </div>
      </q-card-section>

      <!-- georeference-->
      <q-separator />
      <q-card-section v-if="data.georeference">
        <div class="text-h5 text-center">Georeference Actions</div>
        <div class="text-caption text-center">
          These permissions are for managing the specimens georeference data
          fields
        </div>

        <div
          v-for="(value, key) in data.georeference"
          :key="key"
          class="permSection"
        >
          <div class="permName">{{ key }}</div>

          <q-option-group
            inline
            :disable="!isEditable"
            :modelValue="data.georeference[key]"
            :options="RoleOptions"
            color="green"
            type="checkbox"
            @update:modelValue="
              data.georeference = { ...data.georeference, [key]: $event }
            "
          />
        </div>
      </q-card-section>

      <!-- images -->
      <q-separator />
      <q-card-section v-if="data.images">
        <div class="text-h5 text-center">Image Actions</div>
        <div class="text-caption text-center">
          These permissions are for managing the specimens georeference data
          fields
        </div>

        <div v-for="(value, key) in data.images" :key="key" class="permSection">
          <div class="permName">{{ key }}:</div>

          <q-option-group
            inline
            :disable="!isEditable"
            :modelValue="data.images[key]"
            :options="RoleOptions"
            color="green"
            type="checkbox"
            @update:modelValue="data.images = { ...data.images, [key]: $event }"
          />
        </div>
      </q-card-section>

      <!-- loans -->
      <q-separator />
      <q-card-section v-if="data.loans">
        <div class="text-h5 text-center">Loans Actions</div>
        <div class="text-caption text-center">
          These permissions are for managing the specimens loans data fields
        </div>

        <div v-for="(value, key) in data.loans" :key="key" class="permSection">
          <div class="permName">{{ key }}</div>

          <q-option-group
            inline
            :disable="!isEditable"
            :modelValue="data.loans[key]"
            :options="RoleOptions"
            color="green"
            type="checkbox"
            @update:modelValue="data.loans = { ...data.loans, [key]: $event }"
          />
        </div>
      </q-card-section>

      <!-- grants -->
      <q-separator />
      <q-card-section v-if="data.grants">
        <div class="text-h5 text-center">Grants Actions</div>
        <div class="text-caption text-center">
          These permissions are for managing the specimens grants data fields
        </div>

        <div v-for="(value, key) in data.grants" :key="key" class="permSection">
          <div class="permName">{{ key }}</div>

          <q-option-group
            inline
            :disable="!isEditable"
            :modelValue="data.grants[key]"
            :options="RoleOptions"
            color="green"
            type="checkbox"
            @update:modelValue="data.grants = { ...data.grants, [key]: $event }"
          />
        </div>
      </q-card-section>
    </q-card>
  </div>
</template>

<style scoped>
#componentWrapper {
  margin: 0 auto;
}

h4 {
  margin: 0;
  text-align: center;
  text-size-adjust: 90%;
}

.permSection {
  display: inline-grid;
  grid-template-columns: 1fr 4fr;
  width: 100%;
}
.permName {
  align-self: center;
  font-weight: bold;
  justify-self: end;
  text-transform: capitalize;
}

.permOptions {
  align-self: center;
  justify-self: start;
}
</style>

<script setup lang="ts">
import { PlainMessage } from '@bufbuild/protobuf';
import { useVModel } from '@vueuse/core';
import { type } from 'os';

import { schema } from 'src/lib/ccbio';

export interface props {
  accessControl: PlainMessage<schema.Collection_AccessControl>;
  isEditable: boolean;
}

const props = withDefaults(defineProps<props>(), {
  isEditable: false,
});
console.log(schema.Role);

const RoleOptions = Object.keys(schema.Role)
  .filter((key) => isNaN(Number(key)))
  .filter((key) => key !== schema.Role[schema.Role.PUBLIC_UNSPECIFIED])
  .map((key) => {
    return {
      label: key,
      value: schema.Role[key],
    };
  });

console.log(RoleOptions);

const emit = defineEmits(['update:accessControl']);

const data = useVModel(props, 'accessControl', emit);
</script>

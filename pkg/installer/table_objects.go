package installer

var generatedTableCRDV1 = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.6.0-beta.0
  creationTimestamp: null
  name: tables.schemas.schemahero.io
spec:
  group: schemas.schemahero.io
  names:
    kind: Table
    listKind: TableList
    plural: tables
    singular: table
  scope: Namespaced
  versions:
  - name: v1alpha4
    schema:
      openAPIV3Schema:
        description: Table is the Schema for the tables API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: TableSpec defines the desired state of Table
            properties:
              database:
                type: string
              name:
                type: string
              requires:
                items:
                  type: string
                type: array
              schema:
                properties:
                  cassandra:
                    properties:
                      clusteringOrder:
                        properties:
                          column:
                            type: string
                          isDescending:
                            type: boolean
                        required:
                        - column
                        type: object
                      columns:
                        items:
                          properties:
                            isStatic:
                              type: boolean
                            name:
                              type: string
                            type:
                              type: string
                          required:
                          - name
                          - type
                          type: object
                        type: array
                      isDeleted:
                        type: boolean
                      primaryKey:
                        items:
                          items:
                            type: string
                          type: array
                        type: array
                      properties:
                        properties:
                          bloomFilterFPChance:
                            type: string
                          caching:
                            additionalProperties:
                              type: string
                            type: object
                          comment:
                            type: string
                          compaction:
                            additionalProperties:
                              type: string
                            type: object
                          compression:
                            additionalProperties:
                              type: string
                            type: object
                          crcCheckChance:
                            type: string
                          dcLocalReadRepairChance:
                            type: string
                          defaultTTL:
                            type: integer
                          gcGraceSeconds:
                            type: integer
                          maxIndexInterval:
                            type: integer
                          memtableFlushPeriodMs:
                            type: integer
                          minIndexInterval:
                            type: integer
                          readRepairChance:
                            type: string
                          speculativeRetry:
                            type: string
                        type: object
                    type: object
                  cockroachdb:
                    properties:
                      columns:
                        items:
                          properties:
                            attributes:
                              properties:
                                autoIncrement:
                                  type: boolean
                              type: object
                            constraints:
                              properties:
                                notNull:
                                  type: boolean
                              type: object
                            default:
                              type: string
                            name:
                              type: string
                            type:
                              type: string
                          required:
                          - name
                          - type
                          type: object
                        type: array
                      foreignKeys:
                        items:
                          properties:
                            columns:
                              items:
                                type: string
                              type: array
                            name:
                              type: string
                            onDelete:
                              type: string
                            references:
                              properties:
                                columns:
                                  items:
                                    type: string
                                  type: array
                                table:
                                  type: string
                              required:
                              - columns
                              - table
                              type: object
                          required:
                          - columns
                          - references
                          type: object
                        type: array
                      indexes:
                        items:
                          properties:
                            columns:
                              items:
                                type: string
                              type: array
                            isUnique:
                              type: boolean
                            name:
                              type: string
                            type:
                              type: string
                          required:
                          - columns
                          type: object
                        type: array
                      isDeleted:
                        type: boolean
                      json:triggers:
                        items:
                          properties:
                            arguments:
                              items:
                                type: string
                              type: array
                            condition:
                              type: string
                            constraintTrigger:
                              type: boolean
                            events:
                              items:
                                type: string
                              type: array
                            executeProcedure:
                              type: string
                            forEachRun:
                              type: boolean
                            forEachStatement:
                              type: boolean
                            name:
                              type: string
                          required:
                          - events
                          - executeProcedure
                          type: object
                        type: array
                      primaryKey:
                        items:
                          type: string
                        type: array
                    type: object
                  mysql:
                    properties:
                      collation:
                        type: string
                      columns:
                        items:
                          properties:
                            attributes:
                              properties:
                                autoIncrement:
                                  type: boolean
                              type: object
                            charset:
                              type: string
                            collation:
                              type: string
                            constraints:
                              properties:
                                notNull:
                                  type: boolean
                              type: object
                            default:
                              type: string
                            name:
                              type: string
                            type:
                              type: string
                          required:
                          - name
                          - type
                          type: object
                        type: array
                      defaultCharset:
                        type: string
                      foreignKeys:
                        items:
                          properties:
                            columns:
                              items:
                                type: string
                              type: array
                            name:
                              type: string
                            onDelete:
                              type: string
                            references:
                              properties:
                                columns:
                                  items:
                                    type: string
                                  type: array
                                table:
                                  type: string
                              required:
                              - columns
                              - table
                              type: object
                          required:
                          - columns
                          - references
                          type: object
                        type: array
                      indexes:
                        items:
                          properties:
                            columns:
                              items:
                                type: string
                              type: array
                            isUnique:
                              type: boolean
                            name:
                              type: string
                            type:
                              type: string
                          required:
                          - columns
                          type: object
                        type: array
                      isDeleted:
                        type: boolean
                      primaryKey:
                        items:
                          type: string
                        type: array
                    type: object
                  postgres:
                    properties:
                      columns:
                        items:
                          properties:
                            attributes:
                              properties:
                                autoIncrement:
                                  type: boolean
                              type: object
                            constraints:
                              properties:
                                notNull:
                                  type: boolean
                              type: object
                            default:
                              type: string
                            name:
                              type: string
                            type:
                              type: string
                          required:
                          - name
                          - type
                          type: object
                        type: array
                      foreignKeys:
                        items:
                          properties:
                            columns:
                              items:
                                type: string
                              type: array
                            name:
                              type: string
                            onDelete:
                              type: string
                            references:
                              properties:
                                columns:
                                  items:
                                    type: string
                                  type: array
                                table:
                                  type: string
                              required:
                              - columns
                              - table
                              type: object
                          required:
                          - columns
                          - references
                          type: object
                        type: array
                      indexes:
                        items:
                          properties:
                            columns:
                              items:
                                type: string
                              type: array
                            isUnique:
                              type: boolean
                            name:
                              type: string
                            type:
                              type: string
                          required:
                          - columns
                          type: object
                        type: array
                      isDeleted:
                        type: boolean
                      json:triggers:
                        items:
                          properties:
                            arguments:
                              items:
                                type: string
                              type: array
                            condition:
                              type: string
                            constraintTrigger:
                              type: boolean
                            events:
                              items:
                                type: string
                              type: array
                            executeProcedure:
                              type: string
                            forEachRun:
                              type: boolean
                            forEachStatement:
                              type: boolean
                            name:
                              type: string
                          required:
                          - events
                          - executeProcedure
                          type: object
                        type: array
                      primaryKey:
                        items:
                          type: string
                        type: array
                    type: object
                  sqlite:
                    properties:
                      columns:
                        items:
                          properties:
                            attributes:
                              properties:
                                autoIncrement:
                                  type: boolean
                              type: object
                            constraints:
                              properties:
                                notNull:
                                  type: boolean
                              type: object
                            default:
                              type: string
                            name:
                              type: string
                            type:
                              type: string
                          required:
                          - name
                          - type
                          type: object
                        type: array
                      foreignKeys:
                        items:
                          properties:
                            columns:
                              items:
                                type: string
                              type: array
                            name:
                              type: string
                            onDelete:
                              type: string
                            references:
                              properties:
                                columns:
                                  items:
                                    type: string
                                  type: array
                                table:
                                  type: string
                              required:
                              - columns
                              - table
                              type: object
                          required:
                          - columns
                          - references
                          type: object
                        type: array
                      indexes:
                        items:
                          properties:
                            columns:
                              items:
                                type: string
                              type: array
                            isUnique:
                              type: boolean
                            name:
                              type: string
                            type:
                              type: string
                          required:
                          - columns
                          type: object
                        type: array
                      isDeleted:
                        type: boolean
                      primaryKey:
                        items:
                          type: string
                        type: array
                    type: object
                type: object
            required:
            - database
            - name
            type: object
          status:
            description: TableStatus defines the observed state of Table
            properties:
              lastPlannedTableSpecSHA:
                description: We store the SHA of the table spec from the last time
                  we executed a plan to make startup less noisy by skipping re-planning
                  objects that have been planned we cannot use the resourceVersion
                  or generation fields because updating them would cause the object
                  to be modified again
                type: string
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`

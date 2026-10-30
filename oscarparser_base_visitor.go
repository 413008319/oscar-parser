// Code generated from OscarParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // OscarParser
import "github.com/antlr4-go/antlr/v4"

type BaseOscarParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseOscarParserVisitor) VisitSql_script(ctx *Sql_scriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUnit_statement(ctx *Unit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_diskgroup(ctx *Alter_diskgroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_disk_clause(ctx *Add_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_disk_clause(ctx *Drop_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitResize_disk_clause(ctx *Resize_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReplace_disk_clause(ctx *Replace_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWait_nowait(ctx *Wait_nowaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRename_disk_clause(ctx *Rename_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDisk_online_clause(ctx *Disk_online_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDisk_offline_clause(ctx *Disk_offline_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTimeout_clause(ctx *Timeout_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRebalance_diskgroup_clause(ctx *Rebalance_diskgroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPhase(ctx *PhaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCheck_diskgroup_clause(ctx *Check_diskgroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDiskgroup_template_clauses(ctx *Diskgroup_template_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQualified_template_clause(ctx *Qualified_template_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRedundancy_clause(ctx *Redundancy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStriping_clause(ctx *Striping_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitForce_noforce(ctx *Force_noforceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDiskgroup_directory_clauses(ctx *Diskgroup_directory_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDir_name(ctx *Dir_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDiskgroup_alias_clauses(ctx *Diskgroup_alias_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDiskgroup_volume_clauses(ctx *Diskgroup_volume_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_volume_clause(ctx *Add_volume_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_volume_clause(ctx *Modify_volume_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDiskgroup_attributes(ctx *Diskgroup_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_diskgroup_file(ctx *Modify_diskgroup_fileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDisk_region_clause(ctx *Disk_region_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_diskgroup_file_clause(ctx *Drop_diskgroup_file_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConvert_redundancy_clause(ctx *Convert_redundancy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUsergroup_clauses(ctx *Usergroup_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUser_clauses(ctx *User_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFile_permissions_clause(ctx *File_permissions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFile_owner_clause(ctx *File_owner_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitScrub_clause(ctx *Scrub_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQuotagroup_clauses(ctx *Quotagroup_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProperty_name(ctx *Property_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProperty_value(ctx *Property_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFilegroup_clauses(ctx *Filegroup_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_filegroup_clause(ctx *Add_filegroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_filegroup_clause(ctx *Modify_filegroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMove_to_filegroup_clause(ctx *Move_to_filegroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_filegroup_clause(ctx *Drop_filegroup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQuorum_regular(ctx *Quorum_regularContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUndrop_disk_clause(ctx *Undrop_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDiskgroup_availability(ctx *Diskgroup_availabilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEnable_disable_volume(ctx *Enable_disable_volumeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_function(ctx *Drop_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_flashback_archive(ctx *Alter_flashback_archiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_hierarchy(ctx *Alter_hierarchyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_function(ctx *Alter_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_java(ctx *Alter_javaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMatch_string(ctx *Match_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_function_body(ctx *Create_function_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSql_macro_body(ctx *Sql_macro_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitParallel_enable_clause(ctx *Parallel_enable_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartition_by_clause(ctx *Partition_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitResult_cache_clause(ctx *Result_cache_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAccessible_by_clause(ctx *Accessible_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_collation_clause(ctx *Default_collation_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAggregate_clause(ctx *Aggregate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPipelined_using_clause(ctx *Pipelined_using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAccessor(ctx *AccessorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRelies_on_part(ctx *Relies_on_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStreaming_clause(ctx *Streaming_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_outline(ctx *Alter_outlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOutline_options(ctx *Outline_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_lockdown_profile(ctx *Alter_lockdown_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLockdown_feature(ctx *Lockdown_featureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLockdown_options(ctx *Lockdown_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLockdown_statements(ctx *Lockdown_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStatement_clauses(ctx *Statement_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitClause_options(ctx *Clause_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOption_values(ctx *Option_valuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitString_list(ctx *String_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDisable_enable(ctx *Disable_enableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_lockdown_profile(ctx *Drop_lockdown_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_package(ctx *Drop_packageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_package(ctx *Alter_packageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_package(ctx *Create_packageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_package_body(ctx *Create_package_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPackage_obj_spec(ctx *Package_obj_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProcedure_spec(ctx *Procedure_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFunction_spec(ctx *Function_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPackage_obj_body(ctx *Package_obj_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_pmem_filestore(ctx *Alter_pmem_filestoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_pmem_filestore(ctx *Drop_pmem_filestoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_procedure(ctx *Drop_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_procedure(ctx *Alter_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFunction_body(ctx *Function_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProcedure_body(ctx *Procedure_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_procedure_body(ctx *Create_procedure_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_resource_cost(ctx *Alter_resource_costContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_outline(ctx *Drop_outlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_rollback_segment(ctx *Alter_rollback_segmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_restore_point(ctx *Drop_restore_pointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_rollback_segment(ctx *Drop_rollback_segmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_role(ctx *Drop_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_pmem_filestore(ctx *Create_pmem_filestoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPmem_filestore_options(ctx *Pmem_filestore_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFile_path(ctx *File_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_rollback_segment(ctx *Create_rollback_segmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_trigger(ctx *Drop_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_trigger(ctx *Alter_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_trigger(ctx *Create_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTrigger_follows_clause(ctx *Trigger_follows_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTrigger_when_clause(ctx *Trigger_when_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSimple_dml_trigger(ctx *Simple_dml_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFor_each_row(ctx *For_each_rowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCompound_dml_trigger(ctx *Compound_dml_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNon_dml_trigger(ctx *Non_dml_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTrigger_body(ctx *Trigger_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRoutine_clause(ctx *Routine_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCompound_trigger_block(ctx *Compound_trigger_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTiming_point_section(ctx *Timing_point_sectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNon_dml_event(ctx *Non_dml_eventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDml_event_clause(ctx *Dml_event_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDml_event_element(ctx *Dml_event_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDml_event_nested_clause(ctx *Dml_event_nested_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReferencing_clause(ctx *Referencing_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReferencing_element(ctx *Referencing_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_type(ctx *Drop_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_type(ctx *Alter_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCompile_type_clause(ctx *Compile_type_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReplace_type_clause(ctx *Replace_type_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_method_spec(ctx *Alter_method_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_method_element(ctx *Alter_method_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_attribute_definition(ctx *Alter_attribute_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAttribute_definition(ctx *Attribute_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_collection_clauses(ctx *Alter_collection_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDependent_handling_clause(ctx *Dependent_handling_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDependent_exceptions_part(ctx *Dependent_exceptions_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_type(ctx *Create_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitType_definition(ctx *Type_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_type_def(ctx *Object_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_as_part(ctx *Object_as_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_under_part(ctx *Object_under_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNested_table_type_def(ctx *Nested_table_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSqlj_object_type(ctx *Sqlj_object_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitType_body(ctx *Type_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitType_body_elements(ctx *Type_body_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMap_order_func_declaration(ctx *Map_order_func_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubprog_decl_in_type(ctx *Subprog_decl_in_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProc_decl_in_type(ctx *Proc_decl_in_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFunc_decl_in_type(ctx *Func_decl_in_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConstructor_declaration(ctx *Constructor_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModifier_clause(ctx *Modifier_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_member_spec(ctx *Object_member_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSqlj_object_type_attr(ctx *Sqlj_object_type_attrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitElement_spec(ctx *Element_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitElement_spec_options(ctx *Element_spec_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubprogram_spec(ctx *Subprogram_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOverriding_subprogram_spec(ctx *Overriding_subprogram_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOverriding_function_spec(ctx *Overriding_function_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitType_procedure_spec(ctx *Type_procedure_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitType_function_spec(ctx *Type_function_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConstructor_spec(ctx *Constructor_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMap_order_function_spec(ctx *Map_order_function_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPragma_clause(ctx *Pragma_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPragma_elements(ctx *Pragma_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitType_elements_parameter(ctx *Type_elements_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_sequence(ctx *Drop_sequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_sequence(ctx *Alter_sequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_session(ctx *Alter_sessionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_session_set_clause(ctx *Alter_session_set_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_sequence(ctx *Create_sequenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSequence_spec(ctx *Sequence_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSequence_start_clause(ctx *Sequence_start_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_analytic_view(ctx *Create_analytic_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitClassification_clause(ctx *Classification_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCaption_clause(ctx *Caption_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDescription_clause(ctx *Description_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitClassification_item(ctx *Classification_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLanguage(ctx *LanguageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCav_using_clause(ctx *Cav_using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDim_by_clause(ctx *Dim_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDim_key(ctx *Dim_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDim_ref(ctx *Dim_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHier_ref(ctx *Hier_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMeasures_clause(ctx *Measures_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAv_measure(ctx *Av_measureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBase_meas_clause(ctx *Base_meas_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMeas_aggregate_clause(ctx *Meas_aggregate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCalc_meas_clause(ctx *Calc_meas_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_measure_clause(ctx *Default_measure_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_aggregate_clause(ctx *Default_aggregate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCache_clause(ctx *Cache_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCache_specification(ctx *Cache_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLevels_clause(ctx *Levels_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLevel_specification(ctx *Level_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLevel_group_type(ctx *Level_group_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFact_columns_clause(ctx *Fact_columns_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQry_transform_clause(ctx *Qry_transform_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_attribute_dimension(ctx *Create_attribute_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAd_using_clause(ctx *Ad_using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSource_clause(ctx *Source_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJoin_path_clause(ctx *Join_path_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJoin_condition(ctx *Join_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJoin_condition_item(ctx *Join_condition_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAttributes_clause(ctx *Attributes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAd_attributes_clause(ctx *Ad_attributes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAd_level_clause(ctx *Ad_level_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitKey_clause(ctx *Key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlternate_key_clause(ctx *Alternate_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDim_order_clause(ctx *Dim_order_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAll_clause(ctx *All_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_audit_policy(ctx *Create_audit_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPrivilege_audit_clause(ctx *Privilege_audit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAction_audit_clause(ctx *Action_audit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSystem_actions(ctx *System_actionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStandard_actions(ctx *Standard_actionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitActions_clause(ctx *Actions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_action(ctx *Object_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSystem_action(ctx *System_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitComponent_actions(ctx *Component_actionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitComponent_action(ctx *Component_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRole_audit_clause(ctx *Role_audit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_controlfile(ctx *Create_controlfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitControlfile_options(ctx *Controlfile_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLogfile_clause(ctx *Logfile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCharacter_set_clause(ctx *Character_set_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFile_specification(ctx *File_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_diskgroup(ctx *Create_diskgroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQualified_disk_clause(ctx *Qualified_disk_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_edition(ctx *Create_editionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_flashback_archive(ctx *Create_flashback_archiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFlashback_archive_quota(ctx *Flashback_archive_quotaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFlashback_archive_retention(ctx *Flashback_archive_retentionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_hierarchy(ctx *Create_hierarchyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHier_using_clause(ctx *Hier_using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLevel_hier_clause(ctx *Level_hier_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHier_attrs_clause(ctx *Hier_attrs_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHier_attr_clause(ctx *Hier_attr_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHier_attr_name(ctx *Hier_attr_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_index(ctx *Create_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCluster_index_clause(ctx *Cluster_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCluster_name(ctx *Cluster_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_index_clause(ctx *Table_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBitmap_join_index_clause(ctx *Bitmap_join_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_expr(ctx *Index_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_properties(ctx *Index_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDomain_index_clause(ctx *Domain_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLocal_domain_index_clause(ctx *Local_domain_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmlindex_clause(ctx *Xmlindex_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLocal_xmlindex_clause(ctx *Local_xmlindex_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGlobal_partitioned_index(ctx *Global_partitioned_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_partitioning_clause(ctx *Index_partitioning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_partitioning_values_list(ctx *Index_partitioning_values_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLocal_partitioned_index(ctx *Local_partitioned_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOn_range_partitioned_table(ctx *On_range_partitioned_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOn_list_partitioned_table(ctx *On_list_partitioned_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartitioned_table(ctx *Partitioned_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOn_hash_partitioned_table(ctx *On_hash_partitioned_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOn_hash_partitioned_clause(ctx *On_hash_partitioned_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOn_comp_partitioned_table(ctx *On_comp_partitioned_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOn_comp_partitioned_clause(ctx *On_comp_partitioned_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_subpartition_clause(ctx *Index_subpartition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_subpartition_subclause(ctx *Index_subpartition_subclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOdci_parameters(ctx *Odci_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndextype(ctx *IndextypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_index(ctx *Alter_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_index_ops_set1(ctx *Alter_index_ops_set1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_index_ops_set2(ctx *Alter_index_ops_set2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitVisible_or_invisible(ctx *Visible_or_invisibleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMonitoring_nomonitoring(ctx *Monitoring_nomonitoringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRebuild_clause(ctx *Rebuild_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_index_partitioning(ctx *Alter_index_partitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_index_default_attrs(ctx *Modify_index_default_attrsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_hash_index_partition(ctx *Add_hash_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCoalesce_index_partition(ctx *Coalesce_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_index_partition(ctx *Modify_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_index_partitions_ops(ctx *Modify_index_partitions_opsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRename_index_partition(ctx *Rename_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_index_partition(ctx *Drop_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSplit_index_partition(ctx *Split_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_partition_description(ctx *Index_partition_descriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_index_subpartition(ctx *Modify_index_subpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartition_name_old(ctx *Partition_name_oldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNew_partition_name(ctx *New_partition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNew_index_name(ctx *New_index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_inmemory_join_group(ctx *Alter_inmemory_join_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_user(ctx *Create_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_user(ctx *Alter_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_user(ctx *Drop_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_identified_by(ctx *Alter_identified_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIdentified_by(ctx *Identified_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIdentified_other_clause(ctx *Identified_other_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUser_tablespace_clause(ctx *User_tablespace_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQuota_clause(ctx *Quota_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProfile_clause(ctx *Profile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRole_clause(ctx *Role_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUser_default_role_clause(ctx *User_default_role_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPassword_expire_clause(ctx *Password_expire_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUser_lock_clause(ctx *User_lock_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUser_editions_clause(ctx *User_editions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_user_editions_clause(ctx *Alter_user_editions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProxy_clause(ctx *Proxy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitContainer_names(ctx *Container_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSet_container_data(ctx *Set_container_dataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_rem_container_data(ctx *Add_rem_container_dataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitContainer_data_clause(ctx *Container_data_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdminister_key_management(ctx *Administer_key_managementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitKeystore_management_clauses(ctx *Keystore_management_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_keystore(ctx *Create_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOpen_keystore(ctx *Open_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitForce_keystore(ctx *Force_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitClose_keystore(ctx *Close_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBackup_keystore(ctx *Backup_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_keystore_password(ctx *Alter_keystore_passwordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMerge_into_new_keystore(ctx *Merge_into_new_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMerge_into_existing_keystore(ctx *Merge_into_existing_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIsolate_keystore(ctx *Isolate_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUnite_keystore(ctx *Unite_keystoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitKey_management_clauses(ctx *Key_management_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSet_key(ctx *Set_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_key(ctx *Create_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMkid(ctx *MkidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMk(ctx *MkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUse_key(ctx *Use_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSet_key_tag(ctx *Set_key_tagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExport_keys(ctx *Export_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitImport_keys(ctx *Import_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMigrate_keys(ctx *Migrate_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReverse_migrate_keys(ctx *Reverse_migrate_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMove_keys(ctx *Move_keysContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIdentified_by_store(ctx *Identified_by_storeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUsing_algorithm_clause(ctx *Using_algorithm_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUsing_tag_clause(ctx *Using_tag_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSecret_management_clauses(ctx *Secret_management_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_update_secret(ctx *Add_update_secretContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDelete_secret(ctx *Delete_secretContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_update_secret_seps(ctx *Add_update_secret_sepsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDelete_secret_seps(ctx *Delete_secret_sepsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitZero_downtime_software_patching_clauses(ctx *Zero_downtime_software_patching_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWith_backup_clause(ctx *With_backup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIdentified_by_password_clause(ctx *Identified_by_password_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitKeystore_password(ctx *Keystore_passwordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPath(ctx *PathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSecret(ctx *SecretContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAnalyze(ctx *AnalyzeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartition_extention_clause(ctx *Partition_extention_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitValidation_clauses(ctx *Validation_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCompute_clauses(ctx *Compute_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFor_clause(ctx *For_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOnline_or_offline(ctx *Online_or_offlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInto_clause1(ctx *Into_clause1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartition_key_value(ctx *Partition_key_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubpartition_key_value(ctx *Subpartition_key_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAssociate_statistics(ctx *Associate_statisticsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_association(ctx *Column_associationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFunction_association(ctx *Function_associationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndextype_name(ctx *Indextype_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUsing_statistics_type(ctx *Using_statistics_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStatistics_type_name(ctx *Statistics_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_cost_clause(ctx *Default_cost_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCpu_cost(ctx *Cpu_costContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIo_cost(ctx *Io_costContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNetwork_cost(ctx *Network_costContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_selectivity_clause(ctx *Default_selectivity_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_selectivity(ctx *Default_selectivityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStorage_table_clause(ctx *Storage_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUnified_auditing(ctx *Unified_auditingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPolicy_name(ctx *Policy_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAudit_traditional(ctx *Audit_traditionalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAudit_direct_path(ctx *Audit_direct_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAudit_container_clause(ctx *Audit_container_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAudit_operation_clause(ctx *Audit_operation_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAuditing_by_clause(ctx *Auditing_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAudit_user(ctx *Audit_userContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAudit_schema_object_clause(ctx *Audit_schema_object_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSql_operation(ctx *Sql_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAuditing_on_clause(ctx *Auditing_on_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_name(ctx *Model_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_name(ctx *Object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProfile_name(ctx *Profile_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSql_statement_shortcut(ctx *Sql_statement_shortcutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_index(ctx *Drop_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDisassociate_statistics(ctx *Disassociate_statisticsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_indextype(ctx *Drop_indextypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_inmemory_join_group(ctx *Drop_inmemory_join_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFlashback_table(ctx *Flashback_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRestore_point(ctx *Restore_pointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPurge_statement(ctx *Purge_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNoaudit_statement(ctx *Noaudit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRename_object(ctx *Rename_objectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGrant_statement(ctx *Grant_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitContainer_clause(ctx *Container_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRevoke_statement(ctx *Revoke_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRevoke_system_privilege(ctx *Revoke_system_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRevokee_clause(ctx *Revokee_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRevoke_object_privileges(ctx *Revoke_object_privilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOn_object_clause(ctx *On_object_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRevoke_roles_from_programs(ctx *Revoke_roles_from_programsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProgram_unit(ctx *Program_unitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_dimension(ctx *Create_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_directory(ctx *Create_directoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDirectory_name(ctx *Directory_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDirectory_path(ctx *Directory_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_inmemory_join_group(ctx *Create_inmemory_join_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_hierarchy(ctx *Drop_hierarchyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_library(ctx *Alter_libraryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_java(ctx *Drop_javaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_library(ctx *Drop_libraryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_java(ctx *Create_javaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_library(ctx *Create_libraryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPlsql_library_source(ctx *Plsql_library_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCredential_name(ctx *Credential_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLibrary_editionable(ctx *Library_editionableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLibrary_debug(ctx *Library_debugContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCompiler_parameters_clause(ctx *Compiler_parameters_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitParameter_value(ctx *Parameter_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLibrary_name(ctx *Library_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_dimension(ctx *Alter_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLevel_clause(ctx *Level_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHierarchy_clause(ctx *Hierarchy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDimension_join_clause(ctx *Dimension_join_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAttribute_clause(ctx *Attribute_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExtended_attribute_clause(ctx *Extended_attribute_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_one_or_more_sub_clause(ctx *Column_one_or_more_sub_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_view(ctx *Alter_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_view_editionable(ctx *Alter_view_editionableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_view(ctx *Create_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEditioning_clause(ctx *Editioning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitView_options(ctx *View_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitView_alias_constraint(ctx *View_alias_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_view_clause(ctx *Object_view_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInline_constraint(ctx *Inline_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInline_ref_constraint(ctx *Inline_ref_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOut_of_line_ref_constraint(ctx *Out_of_line_ref_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOut_of_line_constraint(ctx *Out_of_line_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConstraint_state(ctx *Constraint_stateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmltype_view_clause(ctx *Xmltype_view_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXml_schema_spec(ctx *Xml_schema_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXml_schema_url(ctx *Xml_schema_urlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitElement(ctx *ElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_tablespace(ctx *Alter_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatafile_tempfile_clauses(ctx *Datafile_tempfile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTablespace_logging_clauses(ctx *Tablespace_logging_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTablespace_group_clause(ctx *Tablespace_group_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTablespace_group_name(ctx *Tablespace_group_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTablespace_state_clauses(ctx *Tablespace_state_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFlashback_mode_clause(ctx *Flashback_mode_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNew_tablespace_name(ctx *New_tablespace_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_tablespace(ctx *Create_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPermanent_tablespace_clause(ctx *Permanent_tablespace_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTablespace_encryption_spec(ctx *Tablespace_encryption_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLogging_clause(ctx *Logging_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExtent_management_clause(ctx *Extent_management_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSegment_management_clause(ctx *Segment_management_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTemporary_tablespace_clause(ctx *Temporary_tablespace_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUndo_tablespace_clause(ctx *Undo_tablespace_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTablespace_retention_clause(ctx *Tablespace_retention_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_tablespace_set(ctx *Create_tablespace_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPermanent_tablespace_attrs(ctx *Permanent_tablespace_attrsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTablespace_encryption_clause(ctx *Tablespace_encryption_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_tablespace_params(ctx *Default_tablespace_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_table_compression(ctx *Default_table_compressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLow_high(ctx *Low_highContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_index_compression(ctx *Default_index_compressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInmmemory_clause(ctx *Inmmemory_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatafile_specification(ctx *Datafile_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTempfile_specification(ctx *Tempfile_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatafile_tempfile_spec(ctx *Datafile_tempfile_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRedo_log_file_spec(ctx *Redo_log_file_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAutoextend_clause(ctx *Autoextend_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMaxsize_clause(ctx *Maxsize_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBuild_clause(ctx *Build_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitParallel_clause(ctx *Parallel_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_materialized_view(ctx *Alter_materialized_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_mv_option1(ctx *Alter_mv_option1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_mv_refresh(ctx *Alter_mv_refreshContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRollback_segment(ctx *Rollback_segmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_mv_column_clause(ctx *Modify_mv_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_materialized_view_log(ctx *Alter_materialized_view_logContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_mv_log_column_clause(ctx *Add_mv_log_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMove_mv_log_clause(ctx *Move_mv_log_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMv_log_augmentation(ctx *Mv_log_augmentationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatetime_expr(ctx *Datetime_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInterval_expr(ctx *Interval_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSynchronous_or_asynchronous(ctx *Synchronous_or_asynchronousContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIncluding_or_excluding(ctx *Including_or_excludingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_materialized_view_log(ctx *Create_materialized_view_logContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNew_values_clause(ctx *New_values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMv_log_purge_clause(ctx *Mv_log_purge_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_materialized_zonemap(ctx *Create_materialized_zonemapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_materialized_zonemap(ctx *Alter_materialized_zonemapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_materialized_zonemap(ctx *Drop_materialized_zonemapContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitZonemap_refresh_clause(ctx *Zonemap_refresh_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitZonemap_attributes(ctx *Zonemap_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitZonemap_name(ctx *Zonemap_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOperator_name(ctx *Operator_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOperator_function_name(ctx *Operator_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_zonemap_on_table(ctx *Create_zonemap_on_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_zonemap_as_subquery(ctx *Create_zonemap_as_subqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_operator(ctx *Alter_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_operator(ctx *Drop_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_operator(ctx *Create_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBinding_clause(ctx *Binding_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_binding_clause(ctx *Add_binding_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitImplementation_clause(ctx *Implementation_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPrimary_operator_list(ctx *Primary_operator_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPrimary_operator_item(ctx *Primary_operator_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOperator_context_clause(ctx *Operator_context_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUsing_function_clause(ctx *Using_function_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_binding_clause(ctx *Drop_binding_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_materialized_view(ctx *Create_materialized_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitScoped_table_ref_constraint(ctx *Scoped_table_ref_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMv_column_alias(ctx *Mv_column_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_mv_refresh(ctx *Create_mv_refreshContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_materialized_view(ctx *Drop_materialized_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_context(ctx *Create_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOracle_namespace(ctx *Oracle_namespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_cluster(ctx *Create_clusterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_profile(ctx *Create_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitResource_parameters(ctx *Resource_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPassword_parameters(ctx *Password_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_lockdown_profile(ctx *Create_lockdown_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStatic_base_profile(ctx *Static_base_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDynamic_base_profile(ctx *Dynamic_base_profileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_outline(ctx *Create_outlineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_restore_point(ctx *Create_restore_pointContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_role(ctx *Create_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_table(ctx *Create_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmltype_table(ctx *Xmltype_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmltype_virtual_columns(ctx *Xmltype_virtual_columnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmltype_column_properties(ctx *Xmltype_column_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmltype_storage(ctx *Xmltype_storageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmlschema_spec(ctx *Xmlschema_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_table(ctx *Object_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_type(ctx *Object_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOid_index_clause(ctx *Oid_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOid_clause(ctx *Oid_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_properties(ctx *Object_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_table_substitution(ctx *Object_table_substitutionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRelational_table(ctx *Relational_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitImmutable_table_clauses(ctx *Immutable_table_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitImmutable_table_no_drop_clause(ctx *Immutable_table_no_drop_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitImmutable_table_no_delete_clause(ctx *Immutable_table_no_delete_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBlockchain_table_clauses(ctx *Blockchain_table_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBlockchain_drop_table_clause(ctx *Blockchain_drop_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBlockchain_row_retention_clause(ctx *Blockchain_row_retention_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBlockchain_hash_and_data_format_clause(ctx *Blockchain_hash_and_data_format_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCollation_name(ctx *Collation_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_properties(ctx *Table_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRead_only_clause(ctx *Read_only_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndexing_clause(ctx *Indexing_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAttribute_clustering_clause(ctx *Attribute_clustering_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitClustering_join(ctx *Clustering_joinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitClustering_join_item(ctx *Clustering_join_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEquijoin_condition(ctx *Equijoin_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCluster_clause(ctx *Cluster_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitClustering_columns(ctx *Clustering_columnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitClustering_column_group(ctx *Clustering_column_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitYes_no(ctx *Yes_noContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitZonemap_clause(ctx *Zonemap_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLogical_replication_clause(ctx *Logical_replication_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRelational_property(ctx *Relational_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_partitioning_clauses(ctx *Table_partitioning_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRange_partitions(ctx *Range_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitList_partitions(ctx *List_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHash_partitions(ctx *Hash_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndividual_hash_partitions(ctx *Individual_hash_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHash_partitions_by_quantity(ctx *Hash_partitions_by_quantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHash_partition_quantity(ctx *Hash_partition_quantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitComposite_range_partitions(ctx *Composite_range_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitComposite_list_partitions(ctx *Composite_list_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitComposite_hash_partitions(ctx *Composite_hash_partitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReference_partitioning(ctx *Reference_partitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReference_partition_desc(ctx *Reference_partition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSystem_partitioning(ctx *System_partitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRange_partition_desc(ctx *Range_partition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitList_partition_desc(ctx *List_partition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubpartition_template(ctx *Subpartition_templateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHash_subpartition_quantity(ctx *Hash_subpartition_quantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubpartition_by_range(ctx *Subpartition_by_rangeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubpartition_by_list(ctx *Subpartition_by_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubpartition_by_hash(ctx *Subpartition_by_hashContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubpartition_name(ctx *Subpartition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRange_subpartition_desc(ctx *Range_subpartition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitList_subpartition_desc(ctx *List_subpartition_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndividual_hash_subparts(ctx *Individual_hash_subpartsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHash_subparts_by_quantity(ctx *Hash_subparts_by_quantityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRange_values_clause(ctx *Range_values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRange_values_list(ctx *Range_values_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitList_values_clause(ctx *List_values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_partition_description(ctx *Table_partition_descriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartitioning_storage_clause(ctx *Partitioning_storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLob_partitioning_storage(ctx *Lob_partitioning_storageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatatype_null_enable(ctx *Datatype_null_enableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSize_clause(ctx *Size_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_compression(ctx *Table_compressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInmemory_table_clause(ctx *Inmemory_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInmemory_attributes(ctx *Inmemory_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInmemory_memcompress(ctx *Inmemory_memcompressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInmemory_priority(ctx *Inmemory_priorityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInmemory_distribute(ctx *Inmemory_distributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInmemory_duplicate(ctx *Inmemory_duplicateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInmemory_column_clause(ctx *Inmemory_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPhysical_attributes_clause(ctx *Physical_attributes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStorage_clause(ctx *Storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDeferred_segment_creation(ctx *Deferred_segment_creationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSegment_attributes_clause(ctx *Segment_attributes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPhysical_properties(ctx *Physical_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIlm_clause(ctx *Ilm_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIlm_policy_clause(ctx *Ilm_policy_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIlm_compression_policy(ctx *Ilm_compression_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIlm_tiering_policy(ctx *Ilm_tiering_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIlm_after_on(ctx *Ilm_after_onContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSegment_group(ctx *Segment_groupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIlm_inmemory_policy(ctx *Ilm_inmemory_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIlm_time_period(ctx *Ilm_time_periodContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHeap_org_table_clause(ctx *Heap_org_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExternal_table_clause(ctx *External_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAccess_driver_type(ctx *Access_driver_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExternal_table_data_props(ctx *External_table_data_propsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOpaque_format_spec(ctx *Opaque_format_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRow_movement_clause(ctx *Row_movement_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFlashback_archive_clause(ctx *Flashback_archive_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLog_grp(ctx *Log_grpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSupplemental_table_logging(ctx *Supplemental_table_loggingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSupplemental_log_grp_clause(ctx *Supplemental_log_grp_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSupplemental_id_key_clause(ctx *Supplemental_id_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAllocate_extent_clause(ctx *Allocate_extent_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDeallocate_unused_clause(ctx *Deallocate_unused_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitShrink_clause(ctx *Shrink_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRecords_per_block_clause(ctx *Records_per_block_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpgrade_table_clause(ctx *Upgrade_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTruncate_table(ctx *Truncate_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_table(ctx *Drop_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_tablespace(ctx *Drop_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_tablespace_set(ctx *Drop_tablespace_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIncluding_contents_clause(ctx *Including_contents_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_view(ctx *Drop_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitComment_on_column(ctx *Comment_on_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEnable_or_disable(ctx *Enable_or_disableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAllow_or_disallow(ctx *Allow_or_disallowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_synonym(ctx *Alter_synonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_synonym(ctx *Create_synonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_synonym(ctx *Drop_synonymContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_spfile(ctx *Create_spfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSpfile_name(ctx *Spfile_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPfile_name(ctx *Pfile_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitComment_on_table(ctx *Comment_on_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitComment_on_materialized(ctx *Comment_on_materializedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_analytic_view(ctx *Alter_analytic_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_add_cache_clause(ctx *Alter_add_cache_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLevels_item(ctx *Levels_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMeasure_list(ctx *Measure_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_drop_cache_clause(ctx *Alter_drop_cache_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_attribute_dimension(ctx *Alter_attribute_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_audit_policy(ctx *Alter_audit_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_cluster(ctx *Alter_clusterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_analytic_view(ctx *Drop_analytic_viewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_attribute_dimension(ctx *Drop_attribute_dimensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_audit_policy(ctx *Drop_audit_policyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_flashback_archive(ctx *Drop_flashback_archiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_cluster(ctx *Drop_clusterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_context(ctx *Drop_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_directory(ctx *Drop_directoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_diskgroup(ctx *Drop_diskgroupContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_edition(ctx *Drop_editionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTruncate_cluster(ctx *Truncate_clusterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCache_or_nocache(ctx *Cache_or_nocacheContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatabase_name(ctx *Database_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_database(ctx *Alter_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatabase_clause(ctx *Database_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStartup_clauses(ctx *Startup_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitResetlogs_or_noresetlogs(ctx *Resetlogs_or_noresetlogsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpgrade_or_downgrade(ctx *Upgrade_or_downgradeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRecovery_clauses(ctx *Recovery_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBegin_or_end(ctx *Begin_or_endContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGeneral_recovery(ctx *General_recoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFull_database_recovery(ctx *Full_database_recoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartial_database_recovery(ctx *Partial_database_recoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartial_database_recovery_10g(ctx *Partial_database_recovery_10gContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitManaged_standby_recovery(ctx *Managed_standby_recoveryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDb_name(ctx *Db_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatabase_file_clauses(ctx *Database_file_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_datafile_clause(ctx *Create_datafile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_datafile_clause(ctx *Alter_datafile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_tempfile_clause(ctx *Alter_tempfile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMove_datafile_clause(ctx *Move_datafile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLogfile_clauses(ctx *Logfile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_logfile_clauses(ctx *Add_logfile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGroup_redo_logfile(ctx *Group_redo_logfileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_logfile_clauses(ctx *Drop_logfile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSwitch_logfile_clause(ctx *Switch_logfile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSupplemental_db_logging(ctx *Supplemental_db_loggingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_or_drop(ctx *Add_or_dropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSupplemental_plsql_clause(ctx *Supplemental_plsql_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLogfile_descriptor(ctx *Logfile_descriptorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitControlfile_clauses(ctx *Controlfile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTrace_file_clause(ctx *Trace_file_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStandby_database_clauses(ctx *Standby_database_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitActivate_standby_db_clause(ctx *Activate_standby_db_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMaximize_standby_db_clause(ctx *Maximize_standby_db_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRegister_logfile_clause(ctx *Register_logfile_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCommit_switchover_clause(ctx *Commit_switchover_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStart_standby_clause(ctx *Start_standby_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStop_standby_clause(ctx *Stop_standby_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConvert_database_clause(ctx *Convert_database_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_settings_clause(ctx *Default_settings_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSet_time_zone_clause(ctx *Set_time_zone_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInstance_clauses(ctx *Instance_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSecurity_clause(ctx *Security_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDomain(ctx *DomainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatabase(ctx *DatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEdition_name(ctx *Edition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFilenumber(ctx *FilenumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFilename(ctx *FilenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPrepare_clause(ctx *Prepare_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_mirror_clause(ctx *Drop_mirror_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLost_write_protection(ctx *Lost_write_protectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCdb_fleet_clauses(ctx *Cdb_fleet_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLead_cdb_clause(ctx *Lead_cdb_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLead_cdb_uri_clause(ctx *Lead_cdb_uri_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProperty_clauses(ctx *Property_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReplay_upgrade_clauses(ctx *Replay_upgrade_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_database_link(ctx *Alter_database_linkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPassword_value(ctx *Password_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLink_authentication(ctx *Link_authenticationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_database(ctx *Create_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatabase_logging_clauses(ctx *Database_logging_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatabase_logging_sub_clause(ctx *Database_logging_sub_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTablespace_clauses(ctx *Tablespace_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEnable_pluggable_database(ctx *Enable_pluggable_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFile_name_convert(ctx *File_name_convertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFilename_convert_sub_clause(ctx *Filename_convert_sub_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTablespace_datafile_clauses(ctx *Tablespace_datafile_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUndo_mode_clause(ctx *Undo_mode_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_tablespace(ctx *Default_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_temp_tablespace(ctx *Default_temp_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUndo_tablespace(ctx *Undo_tablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_database(ctx *Drop_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCreate_database_link(ctx *Create_database_linkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_database_link(ctx *Drop_database_linkContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_tablespace_set(ctx *Alter_tablespace_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_tablespace_attrs(ctx *Alter_tablespace_attrsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_tablespace_encryption(ctx *Alter_tablespace_encryptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTs_file_name_convert(ctx *Ts_file_name_convertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_role(ctx *Alter_roleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRole_identified_clause(ctx *Role_identified_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_table(ctx *Alter_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMemoptimize_read_write_clause(ctx *Memoptimize_read_write_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_table_properties(ctx *Alter_table_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_table_partitioning(ctx *Alter_table_partitioningContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_table_partition(ctx *Add_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_table_partition(ctx *Drop_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMerge_table_partition(ctx *Merge_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_table_partition(ctx *Modify_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSplit_table_partition(ctx *Split_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTruncate_table_partition(ctx *Truncate_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExchange_table_partition(ctx *Exchange_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCoalesce_table_partition(ctx *Coalesce_table_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_interval_partition(ctx *Alter_interval_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartition_extended_names(ctx *Partition_extended_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubpartition_extended_names(ctx *Subpartition_extended_namesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_table_properties_1(ctx *Alter_table_properties_1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_iot_clauses(ctx *Alter_iot_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_mapping_table_clause(ctx *Alter_mapping_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_overflow_clause(ctx *Alter_overflow_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_overflow_clause(ctx *Add_overflow_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpdate_index_clauses(ctx *Update_index_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpdate_global_index_clause(ctx *Update_global_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpdate_all_indexes_clause(ctx *Update_all_indexes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpdate_all_indexes_index_clause(ctx *Update_all_indexes_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpdate_index_partition(ctx *Update_index_partitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpdate_index_subpartition(ctx *Update_index_subpartitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEnable_disable_clause(ctx *Enable_disable_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUsing_index_clause(ctx *Using_index_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_attributes(ctx *Index_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSort_or_nosort(ctx *Sort_or_nosortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExceptions_clause(ctx *Exceptions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMove_table_clause(ctx *Move_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_org_table_clause(ctx *Index_org_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMapping_table_clause(ctx *Mapping_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitKey_compression(ctx *Key_compressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_org_overflow_clause(ctx *Index_org_overflow_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_clauses(ctx *Column_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_collection_retrieval(ctx *Modify_collection_retrievalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCollection_item(ctx *Collection_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRename_column_clause(ctx *Rename_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOld_column_name(ctx *Old_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNew_column_name(ctx *New_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_modify_drop_column_clauses(ctx *Add_modify_drop_column_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_column_clause(ctx *Drop_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_column_clauses(ctx *Modify_column_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_col_properties(ctx *Modify_col_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_col_visibility(ctx *Modify_col_visibilityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_col_substitutable(ctx *Modify_col_substitutableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_column_clause(ctx *Add_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAlter_varray_col_properties(ctx *Alter_varray_col_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitVarray_col_properties(ctx *Varray_col_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitVarray_storage_clause(ctx *Varray_storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLob_segname(ctx *Lob_segnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLob_item(ctx *Lob_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLob_storage_parameters(ctx *Lob_storage_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLob_storage_clause(ctx *Lob_storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_lob_storage_clause(ctx *Modify_lob_storage_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModify_lob_parameters(ctx *Modify_lob_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLob_parameters(ctx *Lob_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLob_deduplicate_clause(ctx *Lob_deduplicate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLob_compression_clause(ctx *Lob_compression_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLob_retention_clause(ctx *Lob_retention_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEncryption_spec(ctx *Encryption_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTablespace(ctx *TablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitVarray_item(ctx *Varray_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_properties(ctx *Column_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLob_partition_storage(ctx *Lob_partition_storageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPeriod_definition(ctx *Period_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStart_time_column(ctx *Start_time_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEnd_time_column(ctx *End_time_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_definition(ctx *Column_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_collation_name(ctx *Column_collation_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIdentity_clause(ctx *Identity_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIdentity_options_parentheses(ctx *Identity_options_parenthesesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIdentity_options(ctx *Identity_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitVirtual_column_definition(ctx *Virtual_column_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAutogenerated_sequence_definition(ctx *Autogenerated_sequence_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEvaluation_edition_clause(ctx *Evaluation_edition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOut_of_line_part_storage(ctx *Out_of_line_part_storageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNested_table_col_properties(ctx *Nested_table_col_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNested_item(ctx *Nested_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubstitutable_column_clause(ctx *Substitutable_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartition_name(ctx *Partition_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSupplemental_logging_props(ctx *Supplemental_logging_propsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_or_attribute(ctx *Column_or_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_type_col_properties(ctx *Object_type_col_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConstraint_clauses(ctx *Constraint_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOld_constraint_name(ctx *Old_constraint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNew_constraint_name(ctx *New_constraint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_constraint_clause(ctx *Drop_constraint_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_constraint(ctx *Add_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_constraint_clause(ctx *Add_constraint_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCheck_constraint(ctx *Check_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDrop_constraint(ctx *Drop_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitEnable_constraint(ctx *Enable_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDisable_constraint(ctx *Disable_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitForeign_key_clause(ctx *Foreign_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReferences_clause(ctx *References_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOn_delete_clause(ctx *On_delete_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUnique_key_clause(ctx *Unique_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPrimary_key_clause(ctx *Primary_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAnonymous_block(ctx *Anonymous_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInvoker_rights_clause(ctx *Invoker_rights_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCall_spec(ctx *Call_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJava_spec(ctx *Java_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitC_spec(ctx *C_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitC_agent_in_clause(ctx *C_agent_in_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitC_parameters_clause(ctx *C_parameters_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitC_external_parameter(ctx *C_external_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitC_property(ctx *C_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDefault_value_part(ctx *Default_value_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSeq_of_declare_specs(ctx *Seq_of_declare_specsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDeclare_spec(ctx *Declare_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitVariable_declaration(ctx *Variable_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubtype_declaration(ctx *Subtype_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCursor_declaration(ctx *Cursor_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitParameter_spec(ctx *Parameter_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitException_declaration(ctx *Exception_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPragma_declaration(ctx *Pragma_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRecord_type_def(ctx *Record_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitField_spec(ctx *Field_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRef_cursor_type_def(ctx *Ref_cursor_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitType_declaration(ctx *Type_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_type_def(ctx *Table_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_indexed_by_part(ctx *Table_indexed_by_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitVarray_type_def(ctx *Varray_type_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSeq_of_statements(ctx *Seq_of_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLabel_declaration(ctx *Label_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSwallow_to_semi(ctx *Swallow_to_semiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAssignment_statement(ctx *Assignment_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitContinue_statement(ctx *Continue_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExit_statement(ctx *Exit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGoto_statement(ctx *Goto_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitElsif_part(ctx *Elsif_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitElse_part(ctx *Else_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLoop_statement(ctx *Loop_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCursor_loop_param(ctx *Cursor_loop_paramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitForall_statement(ctx *Forall_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBounds_clause(ctx *Bounds_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBetween_bound(ctx *Between_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLower_bound(ctx *Lower_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpper_bound(ctx *Upper_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNull_statement(ctx *Null_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRaise_statement(ctx *Raise_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCall_statement(ctx *Call_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPipe_row_statement(ctx *Pipe_row_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSelection_directive(ctx *Selection_directiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitError_directive(ctx *Error_directiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSelection_directive_body(ctx *Selection_directive_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBody(ctx *BodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitException_handler(ctx *Exception_handlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTrigger_block(ctx *Trigger_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTps_block(ctx *Tps_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSql_statement(ctx *Sql_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExecute_immediate(ctx *Execute_immediateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDynamic_returning_clause(ctx *Dynamic_returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitData_manipulation_language_statements(ctx *Data_manipulation_language_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCursor_manipulation_statements(ctx *Cursor_manipulation_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitClose_statement(ctx *Close_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOpen_statement(ctx *Open_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFetch_statement(ctx *Fetch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitVariable_or_collection(ctx *Variable_or_collectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOpen_for_statement(ctx *Open_for_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTransaction_control_statements(ctx *Transaction_control_statementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSet_transaction_command(ctx *Set_transaction_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSet_constraint_command(ctx *Set_constraint_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCommit_statement(ctx *Commit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWrite_clause(ctx *Write_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRollback_statement(ctx *Rollback_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSavepoint_statement(ctx *Savepoint_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCollection_method_call(ctx *Collection_method_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExplain_statement(ctx *Explain_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSelect_only_statement(ctx *Select_only_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSelect_statement(ctx *Select_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWith_clause(ctx *With_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWith_factoring_clause(ctx *With_factoring_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubquery_factoring_clause(ctx *Subquery_factoring_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSearch_clause(ctx *Search_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCycle_clause(ctx *Cycle_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubav_factoring_clause(ctx *Subav_factoring_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubav_clause(ctx *Subav_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHierarchies_clause(ctx *Hierarchies_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFilter_clauses(ctx *Filter_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFilter_clause(ctx *Filter_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_calcs_clause(ctx *Add_calcs_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAdd_calc_meas_clause(ctx *Add_calc_meas_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubquery(ctx *SubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubquery_basic_elements(ctx *Subquery_basic_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubquery_operation_part(ctx *Subquery_operation_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQuery_block(ctx *Query_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSelected_list(ctx *Selected_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSelect_list_elements(ctx *Select_list_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_ref_list(ctx *Table_ref_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_ref(ctx *Table_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_ref_aux(ctx *Table_ref_auxContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_ref_aux_internal_one(ctx *Table_ref_aux_internal_oneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_ref_aux_internal_two(ctx *Table_ref_aux_internal_twoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_ref_aux_internal_thre(ctx *Table_ref_aux_internal_threContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJoin_clause(ctx *Join_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJoin_on_part(ctx *Join_on_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJoin_using_part(ctx *Join_using_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOuter_join_type(ctx *Outer_join_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQuery_partition_clause(ctx *Query_partition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFlashback_query_clause(ctx *Flashback_query_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPivot_clause(ctx *Pivot_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPivot_element(ctx *Pivot_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPivot_for_clause(ctx *Pivot_for_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPivot_in_clause(ctx *Pivot_in_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPivot_in_clause_element(ctx *Pivot_in_clause_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPivot_in_clause_elements(ctx *Pivot_in_clause_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUnpivot_clause(ctx *Unpivot_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUnpivot_in_clause(ctx *Unpivot_in_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUnpivot_in_elements(ctx *Unpivot_in_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHierarchical_query_clause(ctx *Hierarchical_query_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStart_part(ctx *Start_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGroup_by_clause(ctx *Group_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGroup_by_elements(ctx *Group_by_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRollup_cube_clause(ctx *Rollup_cube_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGrouping_sets_clause(ctx *Grouping_sets_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGrouping_sets_elements(ctx *Grouping_sets_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitHaving_clause(ctx *Having_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_clause(ctx *Model_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCell_reference_options(ctx *Cell_reference_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReturn_rows_clause(ctx *Return_rows_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReference_model(ctx *Reference_modelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMain_model(ctx *Main_modelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_column_clauses(ctx *Model_column_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_column_partition_part(ctx *Model_column_partition_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_column_list(ctx *Model_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_column(ctx *Model_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_rules_clause(ctx *Model_rules_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_rules_part(ctx *Model_rules_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_rules_element(ctx *Model_rules_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCell_assignment(ctx *Cell_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_iterate_clause(ctx *Model_iterate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUntil_part(ctx *Until_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOrder_by_clause(ctx *Order_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOrder_by_elements(ctx *Order_by_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOffset_clause(ctx *Offset_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFetch_clause(ctx *Fetch_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFor_update_clause(ctx *For_update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFor_update_of_part(ctx *For_update_of_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFor_update_options(ctx *For_update_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpdate_statement(ctx *Update_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUpdate_set_clause(ctx *Update_set_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_based_update_set_clause(ctx *Column_based_update_set_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDelete_statement(ctx *Delete_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInsert_statement(ctx *Insert_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSingle_table_insert(ctx *Single_table_insertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMulti_table_insert(ctx *Multi_table_insertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMulti_table_element(ctx *Multi_table_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConditional_insert_clause(ctx *Conditional_insert_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConditional_insert_when_part(ctx *Conditional_insert_when_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConditional_insert_else_part(ctx *Conditional_insert_else_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInsert_into_clause(ctx *Insert_into_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitValues_clause(ctx *Values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMerge_statement(ctx *Merge_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMerge_update_clause(ctx *Merge_update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMerge_element(ctx *Merge_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMerge_update_delete_part(ctx *Merge_update_delete_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMerge_insert_clause(ctx *Merge_insert_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSelected_tableview(ctx *Selected_tableviewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLock_table_statement(ctx *Lock_table_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWait_nowait_part(ctx *Wait_nowait_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLock_table_element(ctx *Lock_table_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLock_mode(ctx *Lock_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGeneral_table_ref(ctx *General_table_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStatic_returning_clause(ctx *Static_returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitError_logging_clause(ctx *Error_logging_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitError_logging_into_part(ctx *Error_logging_into_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitError_logging_reject_part(ctx *Error_logging_reject_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDml_table_expression_clause(ctx *Dml_table_expression_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_collection_expression(ctx *Table_collection_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSubquery_restriction_clause(ctx *Subquery_restriction_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSample_clause(ctx *Sample_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSeed_part(ctx *Seed_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExpressions(ctx *ExpressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCursor_expression(ctx *Cursor_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLogical_expression(ctx *Logical_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUnary_logical_expression(ctx *Unary_logical_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUnary_logical_operation(ctx *Unary_logical_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLogical_operation(ctx *Logical_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMultiset_expression(ctx *Multiset_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRelational_expression(ctx *Relational_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCompound_expression(ctx *Compound_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRelational_operator(ctx *Relational_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIn_elements(ctx *In_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBetween_elements(ctx *Between_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConcatenation(ctx *ConcatenationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInterval_expression(ctx *Interval_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_expression(ctx *Model_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitModel_expression_element(ctx *Model_expression_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSingle_column_for_loop(ctx *Single_column_for_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMulti_column_for_loop(ctx *Multi_column_for_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUnary_expression(ctx *Unary_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitImplicit_cursor_expression(ctx *Implicit_cursor_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCollection_expression(ctx *Collection_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCase_statement(ctx *Case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSimple_case_statement(ctx *Simple_case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSimple_case_when_part(ctx *Simple_case_when_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSearched_case_statement(ctx *Searched_case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSearched_case_when_part(ctx *Searched_case_when_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCase_else_part(ctx *Case_else_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAtom(ctx *AtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQuantified_expression(ctx *Quantified_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitString_function(ctx *String_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStandard_function(ctx *Standard_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_function(ctx *Json_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_object_content(ctx *Json_object_contentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_object_entry(ctx *Json_object_entryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_table_clause(ctx *Json_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_array_element(ctx *Json_array_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_on_null_clause(ctx *Json_on_null_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_return_clause(ctx *Json_return_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_transform_op(ctx *Json_transform_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_column_clause(ctx *Json_column_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_column_definition(ctx *Json_column_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_query_returning_clause(ctx *Json_query_returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_query_return_type(ctx *Json_query_return_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_query_wrapper_clause(ctx *Json_query_wrapper_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_query_on_error_clause(ctx *Json_query_on_error_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_query_on_empty_clause(ctx *Json_query_on_empty_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_value_return_clause(ctx *Json_value_return_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_value_return_type(ctx *Json_value_return_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitJson_value_on_mismatch_clause(ctx *Json_value_on_mismatch_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNumeric_function_wrapper(ctx *Numeric_function_wrapperContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNumeric_function(ctx *Numeric_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitListagg_overflow_clause(ctx *Listagg_overflow_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOther_function(ctx *Other_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOver_clause_keyword(ctx *Over_clause_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWithin_or_over_clause_keyword(ctx *Within_or_over_clause_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStandard_prediction_function_keyword(ctx *Standard_prediction_function_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOver_clause(ctx *Over_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWindowing_clause(ctx *Windowing_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWindowing_type(ctx *Windowing_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWindowing_elements(ctx *Windowing_elementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUsing_clause(ctx *Using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUsing_element(ctx *Using_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCollect_order_by_part(ctx *Collect_order_by_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWithin_or_over_part(ctx *Within_or_over_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitString_delimiter(ctx *String_delimiterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCost_matrix_clause(ctx *Cost_matrix_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXml_passing_clause(ctx *Xml_passing_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXml_attributes_clause(ctx *Xml_attributes_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXml_namespaces_clause(ctx *Xml_namespaces_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXml_table_column(ctx *Xml_table_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXml_general_default_part(ctx *Xml_general_default_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXml_multiuse_expression_element(ctx *Xml_multiuse_expression_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmlroot_param_version_part(ctx *Xmlroot_param_version_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmlroot_param_standalone_part(ctx *Xmlroot_param_standalone_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmlserialize_param_enconding_part(ctx *Xmlserialize_param_enconding_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmlserialize_param_version_part(ctx *Xmlserialize_param_version_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmlserialize_param_ident_part(ctx *Xmlserialize_param_ident_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSql_plus_command_no_semicolon(ctx *Sql_plus_command_no_semicolonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSql_plus_command(ctx *Sql_plus_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitStart_command(ctx *Start_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWhenever_command(ctx *Whenever_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSet_command(ctx *Set_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTiming_command(ctx *Timing_commandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPartition_extension_clause(ctx *Partition_extension_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_alias(ctx *Column_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_alias(ctx *Table_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitWhere_clause(ctx *Where_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQuantitative_where_stmt(ctx *Quantitative_where_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInto_clause(ctx *Into_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXml_column_name(ctx *Xml_column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCost_class_name(ctx *Cost_class_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAttribute_name(ctx *Attribute_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSavepoint_name(ctx *Savepoint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRollback_segment_name(ctx *Rollback_segment_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_var_name(ctx *Table_var_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSchema_name(ctx *Schema_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRoutine_name(ctx *Routine_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPackage_name(ctx *Package_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitImplementation_type_name(ctx *Implementation_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitParameter_name(ctx *Parameter_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitReference_model_name(ctx *Reference_model_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitMain_model_name(ctx *Main_model_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitContainer_tableview_name(ctx *Container_tableview_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitAggregate_function_name(ctx *Aggregate_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQuery_name(ctx *Query_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGrantee_name(ctx *Grantee_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRole_name(ctx *Role_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConstraint_name(ctx *Constraint_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLabel_name(ctx *Label_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitType_name(ctx *Type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSequence_name(ctx *Sequence_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitException_name(ctx *Exception_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFunction_name(ctx *Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitProcedure_name(ctx *Procedure_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTrigger_name(ctx *Trigger_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitVariable_name(ctx *Variable_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIndex_name(ctx *Index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCursor_name(ctx *Cursor_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRecord_name(ctx *Record_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitCollection_name(ctx *Collection_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLink_name(ctx *Link_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitLocal_link_name(ctx *Local_link_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConnection_qualifier(ctx *Connection_qualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_name(ctx *Column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTableview_name(ctx *Tableview_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitXmltable(ctx *XmltableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitChar_set_name(ctx *Char_set_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSynonym_name(ctx *Synonym_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSchema_object_name(ctx *Schema_object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDir_object_name(ctx *Dir_object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitUser_object_name(ctx *User_object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGrant_object_name(ctx *Grant_object_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitColumn_list(ctx *Column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitParen_column_list(ctx *Paren_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitKeep_clause(ctx *Keep_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFunction_argument(ctx *Function_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFunction_argument_analytic(ctx *Function_argument_analyticContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitFunction_argument_modeling(ctx *Function_argument_modelingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRespect_or_ignore_nulls(ctx *Respect_or_ignore_nullsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitType_spec(ctx *Type_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitDatatype(ctx *DatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitPrecision_part(ctx *Precision_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNative_datatype_element(ctx *Native_datatype_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitBind_variable(ctx *Bind_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGeneral_element(ctx *General_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitGeneral_element_part(ctx *General_element_partContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitTable_element(ctx *Table_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitObject_privilege(ctx *Object_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitSystem_privilege(ctx *System_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNumeric(ctx *NumericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNumeric_negative(ctx *Numeric_negativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitQuoted_string(ctx *Quoted_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitId_expression(ctx *Id_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitInquiry_directive(ctx *Inquiry_directiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitOuter_join_sign(ctx *Outer_join_signContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitRegular_id(ctx *Regular_idContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNon_reserved_keywords_in_18c(ctx *Non_reserved_keywords_in_18cContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNon_reserved_keywords_in_12c(ctx *Non_reserved_keywords_in_12cContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNon_reserved_keywords_pre12c(ctx *Non_reserved_keywords_pre12cContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitString_function_name(ctx *String_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseOscarParserVisitor) VisitNumeric_function_name(ctx *Numeric_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

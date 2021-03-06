// THIS FILE IS AUTO-GENERATED. DO NOT EDIT
package ai.verta.swagger._public.modeldb.model

import scala.util.Try

import net.liftweb.json._

import ai.verta.swagger._public.modeldb.model.ArtifactTypeEnumArtifactType._
import ai.verta.swagger._public.modeldb.model.AuthzActionEnumAuthzServiceActions._
import ai.verta.swagger._public.modeldb.model.CollaboratorTypeEnumCollaboratorType._
import ai.verta.swagger._public.modeldb.model.DatasetTypeEnumDatasetType._
import ai.verta.swagger._public.modeldb.model.DatasetVisibilityEnumDatasetVisibility._
import ai.verta.swagger._public.modeldb.model.EntitiesEnumEntitiesTypes._
import ai.verta.swagger._public.modeldb.model.IdServiceProviderEnumIdServiceProvider._
import ai.verta.swagger._public.modeldb.model.ModelDBActionEnumModelDBServiceActions._
import ai.verta.swagger._public.modeldb.model.OperatorEnumOperator._
import ai.verta.swagger._public.modeldb.model.PathLocationTypeEnumPathLocationType._
import ai.verta.swagger._public.modeldb.model.RoleActionEnumRoleServiceActions._
import ai.verta.swagger._public.modeldb.model.ServiceEnumService._
import ai.verta.swagger._public.modeldb.model.TernaryEnumTernary._
import ai.verta.swagger._public.modeldb.model.ValueTypeEnumValueType._
import ai.verta.swagger._public.modeldb.model.WorkspaceTypeEnumWorkspaceType._
import ai.verta.swagger._public.modeldb.model.ModeldbProjectVisibility._
import ai.verta.swagger._public.modeldb.model.ProtobufNullValue._
import ai.verta.swagger._public.modeldb.model.UacFlagEnum._
import ai.verta.swagger.client.objects._

case class UacUserInfo (
  dev_key: Option[String] = None,
  email: Option[String] = None,
  first_name: Option[String] = None,
  full_name: Option[String] = None,
  id_service_provider: Option[IdServiceProviderEnumIdServiceProvider] = None,
  image_url: Option[String] = None,
  last_name: Option[String] = None,
  roles: Option[List[String]] = None,
  user_id: Option[String] = None,
  verta_info: Option[UacVertaUserInfo] = None
) extends BaseSwagger {
  def toJson(): JValue = UacUserInfo.toJson(this)
}

object UacUserInfo {
  def toJson(obj: UacUserInfo): JObject = {
    new JObject(
      List[Option[JField]](
        obj.dev_key.map(x => JField("dev_key", JString(x))),
        obj.email.map(x => JField("email", JString(x))),
        obj.first_name.map(x => JField("first_name", JString(x))),
        obj.full_name.map(x => JField("full_name", JString(x))),
        obj.id_service_provider.map(x => JField("id_service_provider", ((x: IdServiceProviderEnumIdServiceProvider) => IdServiceProviderEnumIdServiceProvider.toJson(x))(x))),
        obj.image_url.map(x => JField("image_url", JString(x))),
        obj.last_name.map(x => JField("last_name", JString(x))),
        obj.roles.map(x => JField("roles", ((x: List[String]) => JArray(x.map(JString)))(x))),
        obj.user_id.map(x => JField("user_id", JString(x))),
        obj.verta_info.map(x => JField("verta_info", ((x: UacVertaUserInfo) => UacVertaUserInfo.toJson(x))(x)))
      ).flatMap(x => x match {
        case Some(y) => List(y)
        case None => Nil
      })
    )
  }

  def fromJson(value: JValue): UacUserInfo =
    value match {
      case JObject(fields) => {
        val fieldsMap = fields.map(f => (f.name, f.value)).toMap
        UacUserInfo(
          // TODO: handle required
          dev_key = fieldsMap.get("dev_key").map(JsonConverter.fromJsonString),
          email = fieldsMap.get("email").map(JsonConverter.fromJsonString),
          first_name = fieldsMap.get("first_name").map(JsonConverter.fromJsonString),
          full_name = fieldsMap.get("full_name").map(JsonConverter.fromJsonString),
          id_service_provider = fieldsMap.get("id_service_provider").map(IdServiceProviderEnumIdServiceProvider.fromJson),
          image_url = fieldsMap.get("image_url").map(JsonConverter.fromJsonString),
          last_name = fieldsMap.get("last_name").map(JsonConverter.fromJsonString),
          roles = fieldsMap.get("roles").map((x: JValue) => x match {case JArray(elements) => elements.map(JsonConverter.fromJsonString); case _ => throw new IllegalArgumentException(s"unknown type ${x.getClass.toString}")}),
          user_id = fieldsMap.get("user_id").map(JsonConverter.fromJsonString),
          verta_info = fieldsMap.get("verta_info").map(UacVertaUserInfo.fromJson)
        )
      }
      case _ => throw new IllegalArgumentException(s"unknown type ${value.getClass.toString}")
    }
}

// THIS FILE IS AUTO-GENERATED. DO NOT EDIT
package ai.verta.swagger._public.modeldb.versioning.model

import scala.util.Try

import net.liftweb.json._

import ai.verta.swagger._public.modeldb.versioning.model.ArtifactTypeEnumArtifactType._
import ai.verta.swagger._public.modeldb.versioning.model.DiffStatusEnumDiffStatus._
import ai.verta.swagger._public.modeldb.versioning.model.RepositoryVisibilityEnumRepositoryVisibility._
import ai.verta.swagger._public.modeldb.versioning.model.TernaryEnumTernary._
import ai.verta.swagger._public.modeldb.versioning.model.ValueTypeEnumValueType._
import ai.verta.swagger._public.modeldb.versioning.model.WorkspaceTypeEnumWorkspaceType._
import ai.verta.swagger._public.modeldb.versioning.model.ProtobufNullValue._
import ai.verta.swagger.client.objects._

case class VersioningBlobExpanded (
  blob: Option[VersioningBlob] = None,
  location: Option[List[String]] = None
) extends BaseSwagger {
  def toJson(): JValue = VersioningBlobExpanded.toJson(this)
}

object VersioningBlobExpanded {
  def toJson(obj: VersioningBlobExpanded): JObject = {
    new JObject(
      List[Option[JField]](
        obj.blob.map(x => JField("blob", ((x: VersioningBlob) => VersioningBlob.toJson(x))(x))),
        obj.location.map(x => JField("location", ((x: List[String]) => JArray(x.map(JString)))(x)))
      ).flatMap(x => x match {
        case Some(y) => List(y)
        case None => Nil
      })
    )
  }

  def fromJson(value: JValue): VersioningBlobExpanded =
    value match {
      case JObject(fields) => {
        val fieldsMap = fields.map(f => (f.name, f.value)).toMap
        VersioningBlobExpanded(
          // TODO: handle required
          blob = fieldsMap.get("blob").map(VersioningBlob.fromJson),
          location = fieldsMap.get("location").map((x: JValue) => x match {case JArray(elements) => elements.map(JsonConverter.fromJsonString); case _ => throw new IllegalArgumentException(s"unknown type ${x.getClass.toString}")})
        )
      }
      case _ => throw new IllegalArgumentException(s"unknown type ${value.getClass.toString}")
    }
}

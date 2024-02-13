import { PartialType } from '@nestjs/mapped-types';
import { CreateContainerDto } from './create-container.dto';

export class UpdateContainerDto extends PartialType(CreateContainerDto) {
  name?: string;
  parentContainerId?: number;
}

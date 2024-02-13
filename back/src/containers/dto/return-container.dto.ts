import { Container } from '../entities/container.entity';

export class ReturnContainerDto extends Container {
  parentContainer: ReturnContainerDto;
  childrenContainers: ReturnContainerDto[];
}
